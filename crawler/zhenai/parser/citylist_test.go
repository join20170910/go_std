package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	// contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(any("出错了"))
	}
	//fmt.Printf("%s\n", contents)
	result := ParseCityList(contents)
	const resultSize = 470
	if len(result.Requests) != resultSize {
		t.Errorf("%d, %d", len(result.Requests), resultSize)
	}
	// verify

}
