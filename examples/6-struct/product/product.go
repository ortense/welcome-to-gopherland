package product

type Product struct {
	SKU         string
	Name        string
	Description string
	Price       uint
	Tax         uint
	Available   bool
}

var Joelton = Product{
	SKU:         "1234",
	Name:        "Joelton",
	Description: "Campeão de vendas",
	Price:       9_999_00,
	Tax:         10,
}

func New(sku, name, desc string, price, tax uint) Product {
	return Product{
		SKU:         sku,
		Name:        name,
		Description: desc,
		Price:       price,
		Tax:         tax,
		Available:   false,
	}
}

// p := product.New(sku, name, desc, price, tax)
