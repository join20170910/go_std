package main

import (
	"fmt"
	"time"
)

// 批量取 ch 中的数据
var msg chan int

func main() {
	msg = make(chan int, 2)
	go func(msg chan int) {
		for data := range msg {
			fmt.Println(data)
		}
		fmt.Println(" close ch ")
	}(msg)
	msg <- 11111
	msg <- 22222
	for i := 0; i < 100; i++ {
		msg <- i
	}
	// 关闭 channel
	close(msg)
	time.Sleep(time.Second * 10)
}
