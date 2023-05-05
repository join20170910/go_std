package main

import (
	"bufio"
	"errors"
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
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	//自定义 error
	err = errors.New(" this is a custom error")
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(any(err))
		} else {
			fmt.Printf("%s, %s, %s\n", pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	fmt.Fprintf(writer, "423423424242sfasfasfwaerwerawsfsdf3423423")
	defer writer.Flush()
}
func main() {
	writeFile("fTest.txt")
	//tryDefer()
}
