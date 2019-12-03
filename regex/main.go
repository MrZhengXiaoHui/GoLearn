package main

import (
	"fmt"
	"regexp"
)

const text = `
	My email is ccmouse@gmail.com
	email is abc@163.cn
	email is 111@qq.com
	email is 111@qq.com.cn
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}
}
