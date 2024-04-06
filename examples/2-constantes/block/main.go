package main

import (
	"fmt"
	"time"
)

const (
	label   = "total time"
	message = "hello"
)

func main() {
	startAt := time.Now()

	fmt.Println(message)

	fmt.Println(label, time.Since(startAt))
}
