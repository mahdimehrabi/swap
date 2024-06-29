package currency

type GBP struct {
	*Currency
}

func NewGBP() *GBP {
	gbp := new(GBP)
	gbp.Currency.SetDecimalPlace(2)
	return gbp
}
