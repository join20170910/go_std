package main

import (
	"fmt"
	"sync"
	"time"
)

// 主线程一直在阻塞,直到 wg 减为0 就停止
var wg sync.WaitGroup

func main() {
	// demo01()
	// closed()
	//range01()
	// writeAndReader()
	only()
}

//自读自写
func only() {
	//默认情况下,管道是双向的  可读可写
	// 声明为 只写
	var intChan2 chan<- int
	intChan2 = make(chan int, 3)
	intChan2 <- 20
	fmt.Println("intChan2", intChan2)
	//num := <-intChan2
	//fmt.Println(num)
	// 声明 为 只读
	var intChan3 <-chan int
	if intChan3 != nil {
		num1 := <-intChan3
		fmt.Println("num1:", num1)
	}
}
func writeAndReader() {
	intChan := make(chan int, 50)
	wg.Add(2)
	go writeData(intChan)
	go readData(intChan)
	wg.Wait()
}
func writeData(intChan chan int) {
	defer wg.Done()
	for i := 0; i < 50; i++ {
		intChan <- i
		fmt.Println("写入的数据为:", i)
		time.Sleep(time.Second)
	}
	// 关闭管道
	close(intChan)
}
func readData(intChan chan int) {
	defer wg.Done()
	// 遍历
	for v := range intChan {
		fmt.Println("记取的数据为:", v)
		time.Sleep(time.Second)
	}
}

// 管道遍历
func range01() {

	var intChan chan int
	// 初始化
	intChan = make(chan int, 3)
	intChan <- 10
	num := 20
	intChan <- num
	intChan <- 50

	// 管道 遍历前 没有关闭会出现deadlock 的错误
	close(intChan)
	// 遍历
	for v := range intChan {
		fmt.Printf("%v\n", v)
	}
}

// 关闭功能读取数据测试
func closed() {
	var intChan chan int
	// 初始化
	intChan = make(chan int, 3)
	fmt.Printf("intChan的值: %v\n", intChan)
	fmt.Println("---------------------")
	intChan <- 10
	num := 20
	intChan <- num
	intChan <- 50
	close(intChan)
	fmt.Printf("%v,%v\n", <-intChan, <-intChan)
	// 写入数据 报错
	intChan <- 1000
}
func demo01() {

	// 定义 管道 声明管道 定义一个int 类型的管道
	var intChan chan int
	// 初始化
	intChan = make(chan int, 3)
	fmt.Printf("intChan的值: %v\n", intChan)
	fmt.Println("---------------------")
	intChan <- 10
	num := 20
	intChan <- num
	intChan <- 50
	// 注意不能存放大于容量的数据
	// intChan <- 500
	fmt.Printf("实际长度:%v,容量:%v\n", len(intChan), cap(intChan))
	fmt.Printf("%v,%v\n", <-intChan, <-intChan)
	// 取数 操作 chan 长度 会报错
	fmt.Printf("%v,%v\n", <-intChan, <-intChan)

}
