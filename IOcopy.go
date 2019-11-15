package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

// 创建并写文件
func createFile(i int, name string, dataByte []byte) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	if i == 0 {
		for i := 0; i < 10; i++ {
			file.WriteString(strconv.Itoa(i) + "\n")
		}
	} else {
		file.Write(dataByte)
	}
}

// 读文件
func ReadingFile(name string) (dataByte []byte) {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var buf [100]byte
	var content []byte
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			// 读取结束
			break
		}
		if err != nil {
			fmt.Println("read file err ", err)
			return
		}
		content = append(content, buf[:n]...)
	}
	return content
}

//func main() {
//	// 创建或读取的文件名
//	name := "a1.txt"
//	// 拷贝后的文件名
//	tagetName := "a2.txt"
//	createFile(0, name, nil)       // 创建文件
//	file := ReadingFile(name)      // 读取文件
//	createFile(1, tagetName, file) // 拷贝文件
//}
