package printer

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type Taxable interface {
	SKU() string
	Tax() uint
	Price() uint
	PriceWithTax() uint
}

func ShowPrices(t Taxable) {
	writer := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

	fmt.Fprintln(writer, "---- resumo do pedido ----")

	fmt.Fprintf(writer, "SKU:\t#%s\n", t.SKU())
	fmt.Fprintf(writer, "Preço:\t%.2f\n", float32(t.Price())/100)
	fmt.Fprintf(writer, "Impostos:\t%d%%\n", t.Tax())
	fmt.Fprintf(writer, "Total:\t%d\n", t.PriceWithTax()/100)

	fmt.Fprintln(writer, "------------------------")

	writer.Flush()
}
