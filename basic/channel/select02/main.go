package main

import (
	"fmt"
	"time"
)

func main() {
	// 定义一个int 管道
	intChan := make(chan int, 1)
	go func() {
		time.Sleep(time.Second * 5)
		intChan <- 10
	}()
	//定义一个string 管道
	stringChan := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		stringChan <- "msbgolang"
	}()

	//fmt.Println(<-intChan) // 本身取数据就是阻塞的
	select {
	case v := <-intChan:
		fmt.Println("intChan", v)
	case s := <-stringChan:
		fmt.Println("intChan", s)

	}
}
