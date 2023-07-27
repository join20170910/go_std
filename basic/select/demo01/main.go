package main

import (
	"fmt"
	"time"
)

// select 简单功能
func main() {
	// 定义一个int 管道
	intChan := make(chan int, 1)
	go func() {
		time.Sleep(time.Second * 12)
		intChan <- 12
	}()

	//定义 string 管道
	stringChan := make(chan string, 1)
	go func() {
		stringChan <- "hello world"
	}()

	select {
	case v := <-intChan:
		fmt.Println("intChan:", v)
	case v := <-stringChan:
		fmt.Println("stringChan:", v)
	default:
		fmt.Println("防止select被阻塞")
	}
}
