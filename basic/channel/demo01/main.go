package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for {
		for n := range c {

			fmt.Printf("Worker %d received %c\n", id, n)
		}
	}
}

// chan 做为 发数据
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

// 缓冲区 大小为 3
func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)

}
func closeChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)

}
func main() {
	//chanDemo()
	//bufferedChannel()
	closeChannel()
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
	time.Sleep(time.Millisecond)
}
