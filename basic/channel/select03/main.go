package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
)

// defer + recover 机制处理错误
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go printNum()
	go devide()
	wg.Wait()
}

func devide() {
	defer wg.Done()
	defer func() {
		var r any = recover()
		switch r.(type) {
		case runtime.Error:
			log.Println("运行时错误:", r)
		default:
			log.Println("跳过问题")

		}
	}()
	num1 := 10
	num2 := 0
	result := num1 / num2
	fmt.Printf("%v\n", result)
}

func printNum() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", i)
	}

}
