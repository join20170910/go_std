package main

import "fmt"

type Teacher struct {
	Name   string
	Age    int
	School string
}

func main() {
	//创建老师结构体的实例、对象、变量：
	var t1 Teacher
	fmt.Println(t1)

	t1.Name = "jojo"
	t1.Age = 45
	t1.School = "湖南大学"
	fmt.Println(t1)
	fmt.Println(t1.Age + 10)

}
