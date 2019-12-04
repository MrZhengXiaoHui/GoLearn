package main

import (
	"GoLearn/crawler/engine"
	"GoLearn/crawler/zhenai/parser"
)

func main() {

	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/",
		ParserFunc: parser.ParseCityList,
	})
}
