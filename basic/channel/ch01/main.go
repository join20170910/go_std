package main

import (
	"fmt"
	"time"
)

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * i
	}
	close(out)
}

func consumer(in <-chan int) {
	for num := range in {
		fmt.Printf("num=%d\r\n", num)
	}

}
func main() {
	//默认情况下，是双向的
	// 单向 chan

	// var ch1 chan int 双向的
	//var ch2 chan<- float64 // 单向 只能写数据
	// var ch3 <- chan int // 单向 只能读取 int 数据
	c := make(chan int, 3)
	var send chan<- int = c // send -only

	var read <-chan int = c // recv-only

	send <- 1
	<-read
	d := make(chan int)
	go producer(d)
	go consumer(d)
	time.Sleep(10 * time.Second)

}
