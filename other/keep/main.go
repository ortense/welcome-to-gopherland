package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	startAt := time.Now()
	fmt.Println("Press ctrl+c to end")

	end := make(chan os.Signal, 1)
	signal.Notify(end, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-end:

			fmt.Println("\ntotal time:", time.Since(startAt))
			return
		default:
			fmt.Println("\nuptime:", time.Since(startAt))
			time.Sleep(time.Second)
		}
	}
}
