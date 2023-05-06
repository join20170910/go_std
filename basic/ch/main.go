package main

import (
	"fmt"
	"time"
)

var msg chan string

func main() {
	//有缓冲通道 在有缓冲的通道时候向通道写入一个数据总是 happen before 这个数据被从通道中读取完成
	//对应无缓冲的通道来说从通道接受（获取叫做读取）元素 happen before 向通道发送（写入）数据完成
	// channel 无 缓存 处理方式  GO 有一种 happen-before 机制 保障
	msg = make(chan string, 0)
	go func(msg chan string) {
		data := <-msg
		fmt.Println(data)
	}(msg)
	msg <- "boddy"
	time.Sleep(time.Second * 10)
}
