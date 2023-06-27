package main

import "fmt"

/**
slice 学习
*/

func test01() {
	s := make([]string, 3)
	fmt.Println("emp:", s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)
	l := s[2:5]
	fmt.Println("sl1", l)
	l = s[:len(s)]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}

// 通过数组创建 slice
func test02() {
	var intarr [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("%p", &intarr[0])
	fmt.Println("    -----------------------------")
	fmt.Printf("%p", &intarr[1])
}

// make 函数创建 slice
func test03() {
	intslice := make([]int, 3, 20)
	intslice[0] = 100
	intslice[1] = 110
	intslice[2] = 120
	fmt.Printf("%d %d %d", intslice, len(intslice), cap(intslice))
	fmt.Println("---------------------------------------------")
	for i := 0; i < 100; i++ {
		intslice = append(intslice, i)
	}
	fmt.Printf("%d %d %d", intslice, len(intslice), cap(intslice))
}
func main() {
	test03()
}
