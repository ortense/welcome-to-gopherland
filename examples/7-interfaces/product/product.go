package product

type Product struct {
	sku   string
	price uint
	tax   uint

	Name        string
	Description string
	Available   bool
}

func (p Product) SKU() string {
	return p.sku
}

func (p Product) Price() uint {
	return p.price
}

func (p Product) Tax() uint {
	return p.tax
}

func (p Product) CalculateTax() uint {
	return p.price * p.tax / 100
}

func (p Product) PriceWithTax() uint {
	return p.price + p.CalculateTax()
}

func New(sku, name, desc string, price, tax uint) Product {
	return Product{
		sku:   sku,
		price: price,
		tax:   tax,

		Name:        name,
		Description: desc,
		Available:   false,
	}
}
