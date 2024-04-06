package main

import (
	"exemplos/9-concorrencia/fib"
	"exemplos/9-concorrencia/worker"
	"fmt"
	"sync"
	"time"
)

func main() {
	startAt := time.Now()

	var ch = make(chan int)
	var wg sync.WaitGroup

	go worker.WorkInGroup(1, ch, fib.Fib, &wg)
	go worker.WorkInGroup(2, ch, fib.Fib, &wg)
	go worker.WorkInGroup(3, ch, fib.Fib, &wg)
	go worker.WorkInGroup(4, ch, fib.Fib, &wg)
	go worker.WorkInGroup(5, ch, fib.Fib, &wg)

	for i := range 5 {
		ch <- (i + 45)
	}

	close(ch)

	wg.Wait()

	fmt.Println("total time:", time.Since(startAt))
}
