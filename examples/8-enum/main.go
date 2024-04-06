package main

import (
	"fmt"
	"time"
)

func main() {
	startAt := time.Now()

	fmt.Println("Go não tem enum!")

	fmt.Println("total time:", time.Since(startAt))
}
