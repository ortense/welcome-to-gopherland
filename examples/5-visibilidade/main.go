package main

import (
	"exemplos/5-visibilidade/mypkg"
	"fmt"
	"time"
)

func main() {
	fmt.Println(mypkg.Text)
	fmt.Println("total time:", time.Since(mypkg.Start))
}
