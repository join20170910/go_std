package main

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/zhenai/parser"
	"fmt"
	"regexp"
)

func main() {
	engine.Run(
		engine.Request{Url: "http://www.zhenai.com/zhenghun",
			ParserFunc: parser.ParseCityList})
}

func printCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	// match := re.FindAll(contents, -1)
	matchs := re.FindAllSubmatch(contents, -1)
	for _, m := range matchs {
		fmt.Printf("City: %s, URL: %s\n", m[2], m[1])
		fmt.Println("\n")
	}
	// fmt.Println(match)

}
