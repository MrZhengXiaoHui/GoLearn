package main

func updateSlice(s []int) {
	s[0] = 100
}

//func main() {
//	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
//
//	fmt.Println("arr[2:6] =", arr[2:6]) //arr[2:6] = [2 3 4 5]
//	fmt.Println("arr[:6] =", arr[:6])   //arr[:6] = [0 1 2 3 4 5]
//	s1 := arr[2:]
//	fmt.Println("arr[2:] =", s1) //arr[2:] = [2 3 4 5 6 7]
//	s2 := arr[:]
//	fmt.Println("arr[:] =", s2) //arr[:] = [0 1 2 3 4 5 6 7]
//
//	fmt.Println("After updateSlice(s1)")
//	updateSlice(s1)
//	fmt.Println(s1)  //[100 3 4 5 6 7]
//	fmt.Println(arr) //[0 1 100 3 4 5 6 7]
//
//	fmt.Println("After updateSlice(s2)")
//	updateSlice(s2)
//	fmt.Println(s2)  //[100 1 100 3 4 5 6 7]
//	fmt.Println(arr) //[100 1 100 3 4 5 6 7]
//
//	fmt.Println("Reslice")
//	fmt.Println(s2) //[100 1 100 3 4 5 6 7]
//	s2 = s2[:5]
//	fmt.Println(s2) //[100 1 100 3 4]
//	s2 = s2[2:]
//	fmt.Println(s2) //[100 3 4]
//
//	fmt.Println("Extending slice")
//	arr[0], arr[2] = 0, 2
//	s1 = arr[2:6]
//	s2 = s1[3:5]
//	fmt.Println("s1=", s1)                                              //s1= [2 3 4 5]
//	fmt.Println("s2=", s2)                                              //s2= [5 6]  slice可以向后扩展，不可以向前扩展
//	fmt.Println("arr =", arr)                                           //arr = [0 1 2 3 4 5 6 7]
//	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1)) //s1=[2 3 4 5], len(s1)=4, cap(s1)=6  // cap 最大容量
//	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2)) //s2=[5 6], len(s2)=2, cap(s2)=3  // cap 最大容量
//
//	// 由于值传递的关系，必须接收append的返回值
//	// s2的cap(s2)=3  最后是7  可以替换成10
//	s3 := append(s2, 10)
//	s4 := append(s3, 11)
//	s5 := append(s4, 12)
//	fmt.Println("s3, s4, s5 =", s3, s4, s5) //s3, s4, s5 = [5 6 10] [5 6 10 11] [5 6 10 11 12]
//	// s4 and s5 no longer view arr
//	// 添加元素时如果超过cap，系统会重新分配更大的底层数组
//	fmt.Println("arr =", arr) //arr = [0 1 2 3 4 5 6 10]
//}
