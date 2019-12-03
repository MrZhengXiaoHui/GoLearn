package main

import (
	"fmt"
	"unicode/utf8"
)

/*
	使用range遍历pos，rune对
	使用utf8.RuneCountInString获得字符数量
	使用len获取字节长度
	使用[]byte获得字节

	其他字符串操作 在String库里
*/

func main() {
	s := "测试go语言"
	// 每个中文占三字节 UTF-8
	fmt.Println(s)
	for _, b := range []byte(s) {
		fmt.Printf("%x ", b) // e6 b5 8b e8 af 95 67 6f e8 af ad e8 a8 80
	}
	fmt.Println()

	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d %X) ", i, ch) // (0 6D4B) (3 8BD5) (6 67) (7 6F) (8 8BED) (11 8A00)
	}

	fmt.Println("Rune count:", utf8.RuneCountInString(s)) // Rune count: 6

	bytes := []byte(s) // [230 181 139 232 175 149 103 111 232 175 173 232 168 128]  代表了 ASCII 码的一个字符
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch) // 测 试 g o 语 言
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch) //(0 测) (1 试) (2 g) (3 o) (4 语) (5 言)
	}
	fmt.Println()
}
