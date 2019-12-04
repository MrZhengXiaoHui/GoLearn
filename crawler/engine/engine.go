package engine

import (
	"GoLearn/crawler/fetcher"
	"fmt"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request

	for _,r :=range seeds{
		requests=append(requests,r)
	}
	//将种子拿出队列执行
	for len(requests)>0{
		r :=requests[0]
		requests=requests[1:]
		//爬取页面
		//body html字符串--byte类型
		body,err :=fetcher.Fetch(r.Url)
		fmt.Printf("我正在爬取这个地址Fetching %s\n",r.Url)
		if err!=nil{
			//发生获取错误跳出这次循环
			log.Printf("Fetch error "+"fetching url %s:%v",r.Url,err)
			continue
		}
		//解析
		ParseResult :=r.ParserFunc(body)
		//解析出来的结果里面可能包含种子，需要把种子放进去队列
		//append另外一种形式...表示拼接
		requests =append(requests,ParseResult.Requests...)
		for _,item :=range ParseResult.Items{
			fmt.Printf("爬取所得到到信息：%v\n",item)
		}

	}
}
