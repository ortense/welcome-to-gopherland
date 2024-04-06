package worker

import "fmt"

func Work(id int, ch chan int, task func(int) int) {
	for msg := range ch {
		r := task(msg)
		fmt.Printf("Worker %d: task(%d) = %d\n", id, msg, r)
	}
}
