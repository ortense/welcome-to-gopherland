package main

import (
	"exemplos/9-concorrencia/fib"
	"fmt"
	"time"
)

func main() {
	startAt := time.Now()

	for i := range 5 {
		n := i + 45
		fmt.Printf("fib(%d) %d \n", n, fib.Fib(n))
	}

	fmt.Println("total time:", time.Since(startAt))
}
