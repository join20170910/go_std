package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in   chan int
	done func()
}

func doWorker(id int,
	w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n", id, n)
		w.done()
	}
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWorker(id, w)
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

	wg.Add(20)
	for i, w := range workers {
		w.in <- 'a' + i
		//wg.Add(1)
	}
	for i, w := range workers {
		w.in <- 'A' + i
		//wg.Add(1)
	}
	wg.Wait()
}
