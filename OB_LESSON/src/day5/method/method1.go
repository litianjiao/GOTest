package main

import (
	"fmt"
)

type integer int

//不使用指针只能读（拷贝形参）
func (p integer) print() {
	fmt.Println("p is ", p)
}

//通过指针写形参内容
func (p *integer) set(b integer) {
	*p = b
}

type Student struct {
	Name  string
	Age   int
	Score int
	Sex   int
}

func (p *Student) init(name string, age int, score int) {
	p.Name = name
	p.Age = age
	p.Score = score
	fmt.Println(p)
}
func (p Student) get() Student {
	return p
}
func main() {
	var stu Student
	stu.init("xiaoming", 20, 100)
	stu1 := stu.get()
	fmt.Println(stu1)

	var a integer
	a = 100
	a.print()

	a.set(1000)
	a.print()
}
