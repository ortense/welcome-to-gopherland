package main

import "fmt"

func main() {
	cats := []string{"minato", "hiei"}
	cats = append(cats, "maple")
	l := len(cats)
	c := cap(cats)

	fmt.Println(cats, l, c)
}
