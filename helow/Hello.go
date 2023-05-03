package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Person struct {
	mu     sync.Mutex // 互斥锁
	salary int
	level  int
}

func (p *Person) promote() {
	p.mu.Lock()
	p.salary += 1000
	fmt.Println(p.salary)
	p.level += 1
	fmt.Println(p.level)
	p.mu.Unlock()
}

func do1() {
	do2()
}
func do2() {
	do3()
}

func do3() {
	fmt.Printf("do3 finished ")
}
func add(p *int32) {
	//*p++
	atomic.AddInt32(p, 1)
}

func main() {

	p := Person{level: 1, salary: 1000}

	go p.promote()
	go p.promote()
	go p.promote()

	fmt.Println(p.salary)

	c := int32(0)
	for i := 0; i < 1000; i++ {
		go add(&c)
	}
	time.Sleep(time.Second)
	fmt.Println(c)
}
