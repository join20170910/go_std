package parser

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/model"
	"regexp"
	"strconv"
)

const ageRe = `<td><span class="grayL">年龄：</span>([\d]+)</td>`

func ParseProfile(contents []byte) engine.ParseResult {
	profile := model.Profile{}
	re := regexp.MustCompile(ageRe)
	submatch := re.FindSubmatch(contents)
	if submatch != nil {
		age, err := strconv.Atoi(string(submatch[1]))
		if err != nil {
			profile.Age = age
		}
	}
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}
