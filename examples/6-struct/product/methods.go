package product

func (p Product) calculateTax() uint {
	return p.Price * p.Tax / 100
}

func (p Product) PriceWithTax() uint {
	return p.Price + p.calculateTax()
}
