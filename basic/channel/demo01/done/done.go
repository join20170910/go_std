package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in chan int
	wg *sync.WaitGroup
}

func doWorker(id int, c chan int,
	wg *sync.WaitGroup) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
		wg.Done()
	}
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		wg: wg,
	}
	go doWorker(id, w.in, wg)
	return w
}

func main() {
	chanDemo()
}

func chanDemo() {

	var wg sync.WaitGroup
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	// wg.Add(20)
	for i, w := range workers {
		w.in <- 'a' + i
		wg.Add(1)
	}
	for i, w := range workers {
		w.in <- 'A' + i
		wg.Add(1)
	}
	wg.Wait()
}
