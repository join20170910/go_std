package main

import (
	"fmt"
	"time"
)

// chan 做为 参数传递
func worker(id int, c chan int) {
	for {
		fmt.Printf("Worker %d received %d\n", id, <-c)

	}
}
func main() {
	chanDemo()
}

func chanDemo() {
	c := make(chan int)
	go worker(1, c)

	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)
}
