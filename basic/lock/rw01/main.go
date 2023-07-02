package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.RWMutex
var wg sync.WaitGroup

func main() {
	wg.Add(6)
	for i := 0; i < 5; i++ {
		go read()
	}
	go writer()
	wg.Wait()

}

func read() {
	defer wg.Done()
	lock.RLock()
	fmt.Println("开始读取数据")
	time.Sleep(time.Second)
	fmt.Println("读取数据完成")
	lock.RUnlock()
}

func writer() {
	wg.Done()
	lock.Lock()
	fmt.Println("开始写入数据")
	time.Sleep(time.Second * 10)
	fmt.Println("写入数据完成")
	lock.Unlock()
}
