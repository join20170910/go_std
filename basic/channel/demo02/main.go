package main

import "fmt"

func main() {
	// demo01()
	// closed()
	range01()
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
