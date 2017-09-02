package main

import "fmt"

type integer int //用别名定义新的类型 两种类型不同，不等价
type Student struct {
	Num int
}

type Stu Student //alia

func main() {
	var i integer = 1000
	var j int = 100

	j = int(i)
	fmt.Println(j)

	var a Student
	a = Student{30}
	var b Stu
	a = Student(b)
	fmt.Println(a)
}
