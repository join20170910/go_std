package main

import (
	"fmt"
	"sync"
)

// 锁本质是将并行代码 串行化
func main() {
	var wg sync.WaitGroup
	var rwlock sync.RWMutex
	var num int
	wg.Add(2)
	go func() {

		defer rwlock.Unlock()
		rwlock.Lock()
		num = 12
		wg.Done()
	}()
	go func() {
		rwlock.RLock()
		defer rwlock.Unlock()
		fmt.Println(num)
		wg.Done()
	}()
	wg.Wait()

}
