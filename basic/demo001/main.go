package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	intChan := make(chan int, 100)
	exitChan := make(chan bool, 1)
	go writeData(intChan)
	go readData(intChan, exitChan)
	flag := <-exitChan
	fmt.Println(flag)
}

func writeData(intChan chan int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		tempInt := rand.Intn(4) + 16
		intChan <- tempInt
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		val, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println(val)
	}
	exitChan <- true
	close(exitChan)
}
