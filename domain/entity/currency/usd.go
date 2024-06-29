package currency

import "math/big"

type USD struct {
	*Currency
}

func NewUSD() *USD {
	currency := &Currency{
		I:                  new(big.Int),
		decimalPlaceNumber: 18,
	}
	currency.SetDecimalPlace(2)
	return &USD{
		Currency: currency,
	}
}
