package main

import (
	"exemplos/6-struct/product"
	"fmt"
	"time"
)

func main() {
	startAt := time.Now()

	fmt.Println(product.Joelton)

	fmt.Println("total time:", time.Since(startAt))
}
