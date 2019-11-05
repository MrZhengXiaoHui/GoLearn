package main

type Student struct {
	Username string
	Sex      int
	Score    float32
	Grade    string
}

func NewStudent(usernam string, sex int, score float32, grade string) (stu *Student) {
	stu = &Student{
		Username: usernam,
		Sex:      sex,
		Score:    score,
		Grade:    grade,
	}
	return
}
