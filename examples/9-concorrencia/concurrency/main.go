package main

import (
	"exemplos/9-concorrencia/fib"
	"exemplos/9-concorrencia/worker"
	"fmt"
	"time"
)

func main() {
	startAt := time.Now()

	ch := make(chan int)

	go worker.Work(1, ch, fib.Fib)
	go worker.Work(2, ch, fib.Fib)

	for i := range 5 {
		ch <- (i + 45)
	}

	close(ch)

	fmt.Println("total time:", time.Since(startAt))
}
