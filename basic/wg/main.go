package main

//  类似 java deamon 用法 ， 子 goroutine 运行完了 通知主 goroutine  使用：sync.WaitGroup
// group.Add() 要与 group.Done()配套使用  最后 用 Wait()  它会阻塞住，到 add 为0
import (
	"fmt"
	"sync"
	"time"
)

func main() {
	group := sync.WaitGroup{}
	group.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			time.Sleep(1)
			fmt.Println(i)
			group.Done()
		}(i)
	}
	group.Wait()
	fmt.Println("main finished")
}
