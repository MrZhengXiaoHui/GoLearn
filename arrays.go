package main

import "fmt"

func printArray(arr [5]int)  {
	for i, v := range arr {
		//fmt.Println(arr3[i])
		fmt.Println(i, v)
	}
}

//func main() {
//	var arr1 [5]int
//	arr2 := [3]int{1, 3, 5}
//	arr3 := [...]int{2, 4, 6, 8, 10}
//
//	var grid [4][5]int
//
//	fmt.Println(arr1, arr2, arr3)
//	fmt.Println(grid)
//
//	//for i := 0; i < len(arr3); i++ {
//	//	fmt.Println(arr3[i])
//	//}
//	printArray(arr1)
//	printArray(arr3)
//	//printArray(arr2) 不行arr2是[3] printArray传值是[5]
//	// [10]int 和 [20]int  是不同类型
//}
