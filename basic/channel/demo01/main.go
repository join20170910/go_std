package main

import (
	"fmt"
	"time"
)

// chan 做为 参数传递
func worker(c chan int) {
	for {
		n := <-c
		fmt.Println(n)

	}
}
func main() {
	chanDemo()
}

func chanDemo() {
	c := make(chan int)
	go worker(c)

	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)
}
