package main

import (
	"fmt"
	"time"
)

const (
	a = iota
	b = iota
	c = iota
)

const (
	x = iota + 100
	y = (iota + 2) * 3
	z = iota + x + y
)

func main() {
	startAt := time.Now()

	fmt.Println("a =", a, "b =", b, "c =", c)
	fmt.Println("x =", x, "y =", y, "z =", z)

	fmt.Println("total time:", time.Since(startAt))
}
