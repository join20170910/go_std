package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cpuInfo(ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
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

	//context 包提供了三种函数 withCancel withTimeout, with value
	//
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go cpuInfo(ctx)
	time.Sleep(6 * time.Second)
	cancel()
	wg.Wait()
	fmt.Println("监控完成")

}
