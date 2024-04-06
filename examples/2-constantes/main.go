package main

import (
	"fmt"
	"time"
)

const label = "total time:"

func main() {
	startAt := time.Now()

	fmt.Println(label, time.Since(startAt))
}
