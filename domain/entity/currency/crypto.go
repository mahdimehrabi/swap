package currency

import (
	"math"
	"math/big"
)

type Crypto struct {
	*Currency
}

func NewCrypto() *Crypto {
	return &Crypto{
		Currency: &Currency{
			I:                  new(big.Int),
			decimalPlaceNumber: math.Pow10(18),
		},
	}
}
