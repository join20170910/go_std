package main

import (
	"fmt"
	"time"
)

var number, letter = make(chan bool), make(chan bool)

func printNum() {
	i := 1

	for {
		//接收数据
		<-number
		fmt.Printf("%d%d", i, i+1)
		i += 2
		letter <- true
	}
	// 写入值

}
func printLetter() {
	i := 0
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for {
		<-letter
		if i >= len(str) {
			return
		}
		fmt.Print(str[i : i+2])
		i += 2
		number <- true
	}

}

//交叉打印数字和字母
func main() {
	go printNum()
	go printLetter()
	number <- true
	time.Sleep(time.Second * 100)
}
