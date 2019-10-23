package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	// 计算二进制
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	// 没有初始条件和判断条件  ；号可以省略
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever()  {
	// 死循环
	for {
		fmt.Println("abc")
	}
}

//func main() {
//	fmt.Println(
//		convertToBin(5),
//		convertToBin(13),
//		convertToBin(72387885),
//		convertToBin(0),
//	)
//	printFile("abc.txt")
//	forever()
//}
