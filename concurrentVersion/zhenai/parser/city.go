package parser

import (
	"GoLearn/concurrentVersion/engine"
	"regexp"
)

var (
	cityRe=regexp.MustCompile(`<th><a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^<]+)</a></th>`)
	cityUrlRe =regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte) engine.ParseResult {
	// ^> 非右括号的字符
	//re :=regexp.MustCompile(cityRe)
	matches := cityRe.FindAllSubmatch(contents,-1)
	result :=engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name)
			},
		})
	}
	matches = cityUrlRe.FindAllSubmatch(contents,-1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}
