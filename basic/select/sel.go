package main

import (
	"fmt"
	"time"
)

// select 类似于 switch case语句  selelct 的功能 和操作linux 提供的io
// 的 select .poll epoll
// select 主要作用于多个channel
// 很多时候我们并不会多个goroutine 写同一个channel

// 空结构体
var done = make(chan struct{}) // channel 是多线程安全的 channel 要初始化

func go1(ch1 chan struct{}) {
	time.Sleep(2 * time.Second)
	fmt.Println("执行完成 G1")
	ch1 <- struct{}{}
}

func go2(ch2 chan struct{}) {
	time.Sleep(time.Second)
	fmt.Println("执行完成 G2")
	ch2 <- struct{}{}
}
func main() {
	// 需求：多个 goroutine 都在执行，在主的goroutine中监控，那个执行完成就能立马知道
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	go go1(ch1)
	go go2(ch2)

	select {
	case <-ch1:
		fmt.Println("G1 done")
	case <-ch2:
		fmt.Println("G2 done")
	}
}
