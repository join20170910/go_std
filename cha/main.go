package main

import (
	"fmt"
	"time"
)

/**
  channel 发送 和  接收数据
*/

func watch(p *int) {
	for {
		if *p == 1 {
			fmt.Println("hello")
			break
		}
	}
}

func cwatch(c chan int) {
	if <-c == 1 {
		fmt.Println("hello watch")
	}
}
func main() {

	c1 := make(chan int, 10)
	c2 := make(chan int)
	//非阻塞的Channel 实现
	select {
	case <-c1:
		fmt.Println("c1")
	case c2 <- 1:
		fmt.Println("c2")
	default:
		fmt.Println("none")
	}

	d := make(chan int)
	go cwatch(d)
	time.Sleep(time.Second)
	d <- 1

	c := make(chan string, 5)
	c <- "1" // 发送
	<-c      // 接收
	go func() {
		<-c
	}()

}
