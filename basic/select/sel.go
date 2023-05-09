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

func go1() {
	time.Sleep(time.Second)
	fmt.Println("执行完成 G1")
	done <- struct{}{}
}

func go2() {
	time.Sleep(2 * time.Second)
	time.Sleep(time.Second)
	fmt.Println("执行完成 G2")
	done <- struct{}{}
}
func main() {
	go go1()
	go go2()
	<-done
	<-done
	fmt.Println("done")
}
