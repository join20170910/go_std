package main

import (
	"fmt"
	"time"
)

func main() {
	chanDemo()
}

func chanDemo() {
	c := make(chan int)
	go func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()

	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)
}
