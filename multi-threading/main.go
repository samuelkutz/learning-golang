// THIS IS CRAZY OK IM LOVING IT

package main

import (
	"fmt"
	"time"
)

func worker(workerID int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d \n", workerID, x)

		time.Sleep(time.Second) //work
	}
}

// T0
func main() {
	channel := make(chan int)

	for i := 0; i < 100; i++ {
		go worker(i, channel) // 2kb of memory??? wtf
	}

	for i := 0; i < 1000; i++ {
		channel <- i
	}
}
