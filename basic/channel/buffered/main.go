package main

import (
	"fmt"
	"time"
)

func bufferedChannel() {
	c := make(chan int, 4)
	go worker(0, c)
	//发送数据给C
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	close(c)
	time.Sleep(time.Second)
}

func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

//buffered channel 使用
func main() {
	bufferedChannel()

}
