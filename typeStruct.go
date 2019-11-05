package main

type Student struct {
	name string
}

func (g *Student) getter() string {
	return g.name
}

func (s *Student) setter(name string) {
	s.name = name
}

//func main() {
//	Student1:=Student{name:"zs"}
//	fmt.Println(Student1.getter())
//	Student1.setter("ls")
//	fmt.Println(Student1.getter())
//}
