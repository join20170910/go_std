package main

import (
	"fmt"
)

type worker struct {
	in   chan int
	done chan bool
}

func doWorker(id int, c chan int,
	done chan bool) {
	for {
		for n := range c {
			fmt.Printf("Worker %d received %d\n", id, n)
			done <- true
		}
	}
}

// chan 做为 发数据
func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWorker(id, w.in, w.done)
	return w
}

func main() {
	chanDemo()
}

func chanDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}
	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
		//接收
		<-workers[i].done
	}
	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
		//接收
		<-workers[i].done
	}
}
