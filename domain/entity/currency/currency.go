package currency

import (
	"errors"
	"math"
	"math/big"
)

var StrFormatErr = errors.New("string is not a valid number")

type Currency struct {
	I                  *big.Int `json:"-"`
	decimalPlaceNumber float64  `json:"-"`
}

// SetDecimalPlace for usd is 2 because smaller unit of usd is cent
func (c *Currency) SetDecimalPlace(decimalPlaces int) {
	c.decimalPlaceNumber = math.Pow10(decimalPlaces)
}

func (c *Currency) FromIntString(num string) error {
	_, ok := c.I.SetString(num, 10)
	if !ok {
		return StrFormatErr
	}
	return nil
}

func (c *Currency) ToIntString() string {
	return c.I.String()
}

func (c *Currency) FromFloat(amount float64) {
	f := new(big.Float).SetFloat64(amount)
	f.Mul(f, big.NewFloat(c.decimalPlaceNumber))
	f.Int(c.I)
}

// ToFloat risk of accuracy
func (c *Currency) ToFloat() float64 {
	temp := new(big.Int).Set(c.I)
	f := new(big.Float).SetInt(temp)
	f.Quo(f, big.NewFloat(c.decimalPlaceNumber))
	f64, _ := f.Float64()
	return f64
}

func (c *Currency) ToFloatString() string {
	temp := new(big.Int).Set(c.I)
	f := new(big.Float).SetInt(temp)
	f.Quo(f, big.NewFloat(c.decimalPlaceNumber))
	return f.String()
}

func (c *Currency) FromFloatString(fs string) error {
	f, ok := new(big.Float).SetString(fs)
	if !ok {
		return StrFormatErr
	}
	f.Mul(f, big.NewFloat(c.decimalPlaceNumber))
	f.Int(c.I)
	return nil
}

func (c *Currency) Add(amount float64) {
	f := big.NewFloat(c.ToFloat())
	f.Add(f, big.NewFloat(amount))
	f.Mul(f, big.NewFloat(c.decimalPlaceNumber))
	f.Int(c.I)
}

func (c *Currency) Sub(amount float64) {
	f := big.NewFloat(c.ToFloat())
	f.Sub(f, big.NewFloat(amount))
	f.Mul(f, big.NewFloat(c.decimalPlaceNumber))
	f.Int(c.I)
}

func (c *Currency) Divide(amount float64) {
	f := new(big.Float).SetInt(c.I)
	f.Quo(f, big.NewFloat(c.decimalPlaceNumber))
	f.Quo(f, big.NewFloat(amount))
	f.Mul(f, big.NewFloat(c.decimalPlaceNumber))
	f.Int(c.I)
}

func (c *Currency) Multiply(amount float64) {
	f := big.NewFloat(c.ToFloat())
	f.Mul(f, big.NewFloat(amount))
	f.Mul(f, big.NewFloat(c.decimalPlaceNumber))
	f.Int(c.I)
}
