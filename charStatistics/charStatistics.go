package main

import (
	"fmt"
	"regexp"
)

func CharacterStatistics(str string) {
	//编写一个函数接收一个字符串参数，输出该字符串中包含的中文字符数，英文字符数，数字字符数还有其他字符数。
	quantity := map[string]int{
		"chineseQuantity": 0,
		"englishQuantity": 0,
		"numberQuantity":  0,
		"otherQuantity":   0,
	}
	for _, v := range str {
		if regexp.MustCompile("^[\u4e00-\u9fa5]$").MatchString(string(v)) {
			quantity["chineseQuantity"]++
			continue
		}
		if v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z' {
			quantity["englishQuantity"]++
			continue
		}
		if v >= '0' && v <= '9' {
			quantity["numberQuantity"]++
			continue
		}
		quantity["otherQuantity"]++
	}
	fmt.Println(quantity)
}

func main() {
	CharacterStatistics("11中1aaa一bb-   ")
}
