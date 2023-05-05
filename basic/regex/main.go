package main

import (
	"fmt"
	"regexp"
)

const text = `My email is ccmouse@gmail.com@ABC.COM
EMAIL IS abc@163.com
email2 is  kkk@qq.com.cn
email3 is  john173@qq.com

`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	//match := re.FindAllString(text, -1)
	match := re.FindAllStringSubmatch(text, -1)
	// fmt.Println(match)
	for _, m := range match {
		fmt.Println(m)
	}
}
