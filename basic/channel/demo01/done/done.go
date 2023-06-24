package main

import "fmt"

type worker struct {
	in   chan int
	done chan bool
}

func doWorker(id int, c chan int,
	done chan bool) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
		go func() {
			done <- true
		}()
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

	for i, w := range workers {
		w.in <- 'a' + i

	}
	for i, w := range workers {
		w.in <- 'A' + i
	}
	for _, w := range workers {
		<-w.done
		<-w.done
	}
}
