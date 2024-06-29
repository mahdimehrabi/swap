package currency

type EURO struct {
	*Currency
}

func NewEURO() *EURO {
	euro := new(EURO)
	euro.Currency.SetDecimalPlace(2)
	return euro
}
