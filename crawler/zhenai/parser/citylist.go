package parser

import (
	"awesomeProject/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {

	re := regexp.MustCompile(cityListRe)
	matchs := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	limit := 10
	for _, m := range matchs {
		result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{Url: string(m[1]),
			ParserFunc: ParseCity})
		//fmt.Printf("City: %s, URL: %s\n", m[2], m[1])
		//fmt.Println("\n")
		limit--
		if limit <= 0 {
			break
		}
	}
	return result
}
