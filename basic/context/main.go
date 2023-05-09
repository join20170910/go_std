package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cpuInfo(stop chan struct{}) {
	defer wg.Done()
	for {
		select {
		case <-stop:
			fmt.Println("退出cpu监控")
			return
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("cpu 的信息")
		}
	}

}
func main() {
	// 有一个 goroutine  监控 cpu 信息
	var stop = make(chan struct{})

	wg.Add(1)
	go cpuInfo(stop)
	time.Sleep(6 * time.Second)
	stop <- struct{}{}
	wg.Wait()
	fmt.Println("监控完成")

}
