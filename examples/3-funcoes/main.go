package main

import "fmt"

func soma(a int, b int) int {
	return a + b
}

func divisao(a, b int) (int, int) {
	resultado := a / b
	modulo := a % b

	return resultado, modulo
}

func main() {
	a := soma(4, 4)
	b, c := divisao(10, 3)

	fmt.Println(a, b, c)
}
