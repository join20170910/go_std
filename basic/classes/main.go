package main

import "fmt"

type Teacher struct {
	Name   string
	Age    int
	School string
}

func method01() {
	var t1 Teacher
	fmt.Println(t1)

	t1.Name = "jojo"
	t1.Age = 45
	t1.School = "湖南大学"
	fmt.Println(t1)
	fmt.Println(t1.Age + 10)
}
func method02() {
	var t Teacher = Teacher{"john", 35, "cs"}
	fmt.Println(t)
}

// 返回结构体指针
func method03() {
	var t *Teacher = new(Teacher)
	(*t).Name = "张三"
	(*t).Age = 40
	(*t).School = "BD"
	fmt.Println(*t)
}
func method04() {
	var t *Teacher = &Teacher{"ls", 24, "qh"}
	fmt.Println(*t)
}

// 结构体的方法引用  结构体类型是值类型，在方法调用中，遵守值类型的传递机制，是值拷贝传递方式
func (t Teacher) say() {
	fmt.Println(t.Name)
}
func main() {
	//创建老师结构体的实例、对象、变量：
	//方式1
	//method01()
	//method02()
	//method03()
	// method04()
	// var t *Teacher = &Teacher{"ZS", 18, "BD"}
	var t Teacher
	t.Name = "LILI"
	t.say()
	t.Name = "dd"
	fmt.Println(t.Name)

	t.say()

}
