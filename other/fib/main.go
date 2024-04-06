package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	startAt := time.Now()

	for i := range 5 {
		n := i + 45
		fmt.Printf("fib(%d) %d \n", n, fib(n))
	}

	fmt.Println("total time:", time.Since(startAt))
}

func fib(n int) *big.Int {
	if n <= 1 {
		return big.NewInt(int64(n))
	}

	prev1 := big.NewInt(1)
	prev2 := big.NewInt(0)
	current := big.NewInt(0)

	for i := 2; i <= n; i++ {
		current.Add(prev1, prev2)
		prev2.Set(prev1)
		prev1.Set(current)
	}

	return current
}
