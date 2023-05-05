package main

import (
	"bufio"
	"fmt"
	"os"
)

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic(any("error occurred"))
	return
	fmt.Println(4)
}
func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		//panic(any(err))
		fmt.Println("Error:", err.Error())
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	fmt.Fprintf(writer, "423423424242sfasfasfwaerwerawsfsdf3423423")
	defer writer.Flush()
}
func main() {
	writeFile("fTest.txt")
	tryDefer()
}
