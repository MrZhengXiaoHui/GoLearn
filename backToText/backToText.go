package main

import "fmt"

func BackToText(str string) {
	//编写一个函数接受一个字符串参数，输出该字符串是否是回文（从前往后和从后往前都是一样的字符串成为回文）
	str1 := ""
	for i := len(str) - 1; i >= 0; i-- {
		str1 += string(str[i])
	}
	fmt.Println(str == str1)
}

func main() {
	BackToText("abccba") //true
	BackToText("abcbac") //false
}
