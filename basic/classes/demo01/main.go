package main

import "fmt"

// Student 结构体进行type 重新定义
type Student struct {
	Age int
}
type Person struct {
	Name string
}

func (p Person) disp() {
	fmt.Println(p.Name)
}

type Stu Student

func (a Student) test() {
	fmt.Println(a.Age)
}
func main() {
	var s1 Student = Student{Age: 19}
	var s2 Stu = Stu{Age: 19}
	// s1 = Student{s2}
	fmt.Println(s1.Age)
	fmt.Println(s2)
	s1.test()
	var p Person
	p.Name = "lisa"
	p.disp()
}
