package currency

import (
	"math/big"
	"testing"
)

// tests don't pass because we have lack of some features but precision safety is ok
func TestCurrency_Cast(t *testing.T) {
	criticalNumber := 3.3333

	tests := []struct {
		input        float64
		decimalPlace int
		multipleRes  string // multiple result of input by criticalNumber
		sumRes       string // sum result of input by criticalNumber
		divideRes    string // divide result of input by criticalNumber
		subRes       string // sub result of input by criticalNumber
	}{
		{1234.56, 7, "4115.158848", "1237.8933", "370.3717037", "1231.2267"},
		{0.01, 4, "0.0333", "3.3433", "0.003", "-3.3233"},
		{1000000.00, 4, "3333300", "1000003.3333", "300003.0000", "999996.6667"},
		{0.0001, 18, "0.00033333", "3.3334", "0.000030000300003000", "-3.3332"},
		{1, 18, "3.3333", "4.3333", "0.300003000030000300", "-2.3333"},
		{9999999999.99, 10, "33332999999.966667", "10000000003.3233", "3000030000.2970029700", "9999999996.6567"},
		{333333.3333333333333, 18, "1111099.99999999999988889", "333336.6666333333", "100001.000010000099990999",
			"333330.0000333333333"},
		{0.123456, 6, "0.411515", "3.456756", "0.0370371", "-3.209844"},
	}

	for _, test := range tests {
		c := &Currency{I: new(big.Int)}
		c.SetDecimalPlace(test.decimalPlace)
		c.FromFloat(test.input)
		f := c.ToFloat()
		if f != test.input {
			t.Errorf("FromFloat(%f) with decimalPlace %d and ToFloat return are not the same: got %f",
				test.input, test.decimalPlace, f)
		}

		// Test string conversion
		str := c.ToIntString()
		c = &Currency{I: new(big.Int)}
		c.SetDecimalPlace(test.decimalPlace)
		err := c.FromIntString(str)
		if err != nil {
			t.Errorf("FromIntString(%s) with decimalPlace %d returned error: %s", str, test.decimalPlace, err.Error())
		}
		f = c.ToFloat()
		if f != test.input {
			t.Errorf("ToFloat after FromIntString(%s) with decimalPlace %d failed: got %f, want %f",
				str, test.decimalPlace, f, test.input)
		}

		// Test multiplication
		c.FromFloat(test.input)
		c.Multiply(criticalNumber)
		if c.ToFloatString() != test.multipleRes {
			t.Errorf("Multiply(%f) to %f decimalPlace %d failed: got %s, want %s",
				criticalNumber, test.input, test.decimalPlace, c.ToFloatString(), test.multipleRes)
		}

		// Test addition
		c.FromFloat(test.input)
		c.Add(criticalNumber)
		if c.ToFloatString() != test.sumRes {
			t.Errorf("Add(%f) to %f decimalPlace %d failed: got %s, want %s",
				criticalNumber, test.input, test.decimalPlace, c.ToFloatString(), test.sumRes)
		}

		// Test division
		c.FromFloat(test.input)
		c.Divide(criticalNumber)
		if c.ToFloatString() != test.divideRes {
			t.Errorf("Divide(%f) to %f decimalPlace %d failed: got %s, want %s",
				criticalNumber, test.input, test.decimalPlace, c.ToFloatString(), test.divideRes)
		}

		// Test subtraction
		c.FromFloat(test.input)
		c.Sub(criticalNumber)
		f = c.ToFloat()
		if c.ToFloatString() != test.subRes {
			t.Errorf("Sub(%f) to %f decimalPlace %d failed: got %s, want %s",
				criticalNumber, test.input, test.decimalPlace, c.ToFloatString(), test.subRes)
		}

	}
}

func TestCurrency_FromString(t *testing.T) {
	tests := []struct {
		input    string
		expected string // expectedStr value in smaller units
		valid    bool
	}{
		{"123456", "123456", true},
		{"0", "0", true},
		{"100000000", "100000000", true},
		{"invalid", "", false},
		{"123.45", "", false}, // invalid because it should be an integer string
	}

	for _, test := range tests {
		c := &Currency{I: new(big.Int)}
		err := c.FromIntString(test.input)

		if (err == nil) != test.valid {
			t.Errorf("FromIntString(%s) valid = %v; expectedStr %v", test.input, err == nil, test.valid)
		}

		if err == nil && c.ToIntString() != test.expected {
			t.Errorf("FromIntString(%s) = %s; expectedStr %s", test.input, c.ToIntString(), test.expected)
		}
	}
}
