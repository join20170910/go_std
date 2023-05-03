package main

import (
	"fmt"
	"regexp"

	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(any(err))
	}

	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		all, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(any(err))
		}

		// fmt.Printf("%s\n", all)
		printCityList(all)
	} else {
		fmt.Println("Error: status code", resp.StatusCode)
	}
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
