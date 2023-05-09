package main

func main() {
	//默认情况下，是双向的
	// 单向 chan

	// var ch1 chan int 双向的
	//var ch2 chan<- float64 // 单向 只能写数据
	// var ch3 <- chan int // 单向 只能读取 int 数据
	c := make(chan int, 3)
	var send chan<- int = c // send -only

	var read <-chan int = c // recv-only

	send <- 1
	<-read

}
