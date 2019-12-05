package main

import (
	"GoLearn/concurrentVersion/engine"
	"GoLearn/concurrentVersion/persist"
	"GoLearn/concurrentVersion/scheduler"
	"GoLearn/concurrentVersion/zhenai/parser"
)

func main() {

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:persist.ItemSaver(),
	}
	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun/",
	//	ParserFunc: parser.ParseCityList,
	//})
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCity,
	})
}
