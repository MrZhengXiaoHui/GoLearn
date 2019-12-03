package main

import "fmt"

func main() {
	//9*9乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d ", j, i, i*j)
		}
		fmt.Println("")
	}
}
