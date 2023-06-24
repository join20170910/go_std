package main

import (
	"fmt"
	"time"
)

// chan 做为 参数传递
func worker(id int, c chan int) {
	for {
		fmt.Printf("Worker %d received %c\n", id, <-c)

	}
}
func main() {
	chanDemo()
}

func chanDemo() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go worker(i, channels[i])
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}
