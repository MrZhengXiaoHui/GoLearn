package main

import (
	"fmt"
)

var (
	AllStudents []*Student
)

//循环打印帮助信息
func showMeu() {
	fmt.Println("1、添加学生")
	fmt.Println("2、修改学生信息")
	fmt.Println("3、打印所以学生列表")
	fmt.Println("4、结束程序")
}

// 添加学生的方法
func AddStudent() {
	stu := inputStudent()
	// 假设没有重名的学生，名字若一样，则更新
	for index, v := range AllStudents {
		if v.Username == stu.Username {
			// 做更新
			AllStudents[index]=stu
			fmt.Println("更新成功")
			return
		}
	}
	// 不重名直接加
	AllStudents = append(AllStudents, stu)
	fmt.Println("学生插入成功")
}

// 用户输入的方法
func inputStudent() *Student {
	var (
		name  string
		sex   int
		grade string
		score float32
	)
	fmt.Println("请输入学生姓名")
	_, _ = fmt.Scan(&name)
	fmt.Println("请输入学生性别[0|1]")
	_, _ = fmt.Scan(&sex)
	fmt.Println("请输入年级[0-6]")
	_, _ = fmt.Scan(&grade)
	fmt.Println("请输入分数[0-100]")
	_, _ = fmt.Scan(&score)
	stu := NewStudent(name, sex, score, grade)
	return stu
}

func ShowAllStudent(){
	for _,v:=range AllStudents{
		fmt.Printf("学生：%s 信息：%#v\n",v.Username,v)
	}
	fmt.Println()
}

//func main() {
//	for {
//		showMeu()
//
//		var i int
//		fmt.Scan(&i)
//		switch i {
//		case 1:
//			AddStudent()
//		case 3:
//			ShowAllStudent()
//		case 4:
//			os.Exit(0)
//		}
//	}
//}
