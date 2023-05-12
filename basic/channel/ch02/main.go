package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for {
		fmt.Printf("Worker %d received %c\n", id, <-c)
	}
}

// channel 作为返回值  发送数据: chan<- int 接收数据: <-chan int
func createWorker(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("Worker %d received %c\n", id, <-c)
		}
	}()
	return c
}
func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i

	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	// 发送数据
	time.Sleep(time.Microsecond)
}
func main() {
	chanDemo()
}
