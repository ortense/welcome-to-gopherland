package main

import (
	"fmt"
	"time"
)

func main() {
	var start time.Time = time.Now()
	var since = time.Since(start)
	label := "total time"

	fmt.Println(label, since)
}
