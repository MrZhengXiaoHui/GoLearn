package parser

import (
	"GoLearn/crawler/engine"
	"regexp"
)

const cityRe=`<th><a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^<]+)</a></th>`

func ParseCity(contents []byte) engine.ParseResult {
	// ^> 非右括号的字符
	re :=regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents,-1)
	result :=engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items, "User :"+name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name)
			},
		})
	}

	return result
}
