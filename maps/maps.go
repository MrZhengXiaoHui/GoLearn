package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	m2 := make(map[string]int) // m2 == empty map
	var m3 map[string]int      // m3 == nil
	fmt.Println(m, m2, m3)     // map[course:golang name:ccmouse quality:notbad site:imooc] map[]  map[]

	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v) // map 是的key无序的 所以range遍历 map是无序输出的
	}

	fmt.Println("Getting values")
	//courseName := m["course"]
	//fmt.Println(courseName) // golang
	//causeName := m["cause"] // cause不存在这个key值
	//fmt.Println(causeName) // 空串
	// 判断是否存在
	courseName, ok := m["course"]
	fmt.Println(courseName, ok) // golang true
	//causeName, ok := m["cause"] // cause不存在这个key值
	//fmt.Println(causeName, ok)  // 空串 false
	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok) // ccmouse true
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok) // 空串 false
}
