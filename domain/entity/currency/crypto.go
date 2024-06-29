package currency

import "math/big"

type Crypto struct {
	*Currency
}

func NewCrypto() *Crypto {
	return &Crypto{
		Currency: &Currency{
			I:                  new(big.Int),
			decimalPlaceNumber: 18,
		},
	}
}
