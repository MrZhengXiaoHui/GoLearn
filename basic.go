package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//var aa = 3
//var ss = "kk"

//在函数外部定义不可以直接用：号，要有var关键字
//bb := true

var (
	aa = 3
	ss = "kk"
)

func variableZeroValue()  {
	var a int
	var s string

	fmt.Printf("%d %q\n", a, s)

}

func variableInitialValue()  {
	var a, b int = 3, 4
	var s string = "abc"

	fmt.Println(a, b, s)
}

func variableTypeDeduction()  {
	var a, b, c, s = 3, 4, true, "abc"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "abc"
	b = 5
	fmt.Println(a, b, c, s)
}

func euler()  {
	// 复数
	//c := 3 + 4i
	//fmt.Println(cmplx.Abs(c))
	fmt.Printf("%.3f\n",cmplx.Exp(1i * math.Pi) + 1)
}

func triangle()  {
	// 强制类型转换
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(c)
}

func consts() {
	// 常量
	//const filename = "abc.txt"
	//const a, b  = 3, 4
	const(
		filename = "abc.txt"
		a, b  = 3, 4
	)
	var c int
	// 常量的数值可作为各种类型使用
	c = int(math.Sqrt(a * a + b * b))
	fmt.Println(filename, c)
}

func enums() {
	// 枚举
	//const(
	//	cpp = 0
	//	java = 1
	//	python = 2
	//	golang = 3
	//)

	// 简化  iota自增值
	//const(
	//	cpp = iota
	//	java
	//	python
	//	golang
	//)
	//fmt.Println(cpp, java, python, golang) // 0 1 2 3

	// _ 占位
	const(
		cpp = iota
		_
		python
		golang
		javascript
	)
	fmt.Println(cpp, python, golang, javascript) // 0 2 3 4

	// 参与运算 写表达式
	const(
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

//func main()  {
//	fmt.Println("hello go")
//	variableZeroValue()
//	variableInitialValue()
//	variableTypeDeduction()
//	variableShorter()
//	fmt.Println(aa, ss)
//
//	euler()
//	triangle()
//	consts()
//	enums()
//}
