package currency

type IRR struct {
	*Currency
}

func NewIRR() *IRR {
	irr := new(IRR)
	irr.Currency.SetDecimalPlace(0)
	return irr
}
