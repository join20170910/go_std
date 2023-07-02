// channel + goroutine 结合应用
package main

import (
	"fmt"
	"time"
)

var intChan chan int = make(chan int, 100)

func main() {
	start := time.Now()
	var isPrimChan chan int = make(chan int, 100)
	var exitChan chan bool = make(chan bool, 100)
	//初始化
	go initChan(100)
	// 判断素数
	for i := 0; i < 100; i++ {
		go isPrime(intChan, isPrimChan, exitChan)
	}
	go func() {
		for i := 0; i < 10; i++ {
			<-exitChan
		}
		close(isPrimChan)
	}()

	end := time.Since(start)

	for {
		res, ok := <-isPrimChan
		if !ok {
			break
		}
		fmt.Printf("素数: %d\n", res)
	}

	fmt.Println(end)

}

// 初始化 数字放入管道
func initChan(num int) {
	for i := 1; i < num; i++ {
		intChan <- i
	}
	close(intChan)
}

func isPrime(intChan chan int, isPrimChan chan int, exitChan chan bool) {
	var flag bool
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}

		}
		if flag {
			isPrimChan <- num
		}
	}
	exitChan <- true
}
