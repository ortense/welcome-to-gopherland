package worker

import (
	"fmt"
	"sync"
)

func WorkInGroup(id int, ch chan int, task func(int) int, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	for msg := range ch {
		r := task(msg)
		fmt.Printf("Worker %d: task(%d) = %d\n", id, msg, r)
	}
}
