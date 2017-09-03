package main

import (
	"fmt"
)

type People struct {
	name string
	age  int
}

type Act interface { //一个对象如果要实现此接口
	Print()
	Sleep() //需要分别实现这两个方法
}

type Student struct {
	name  string
	age   int
	score int
}

func (p Student) Print() {
	fmt.Println("name:", p.name)
	fmt.Println("age:", p.age)
	fmt.Println("score:", p.score)
}

func (p Student) Sleep() {
	fmt.Println("student sleep")
}

func (pp People) Print() {
	fmt.Println("name:", pp.name)
	fmt.Println("age:", pp.age)
}
func (p People) Sleep() {
	fmt.Println("people sleep")
}

func main() {
	var t Act
	fmt.Println(t)
	var stu Student = Student{
		name:  "xiaoming",
		age:   18,
		score: 100,
	}
	t = stu
	t.Print()
	t.Sleep()

	var people People = People{
		name: "people",
		age:  100,
	}
	t = people
	t.Print()
	t.Sleep()
	//多态的实现 t既可以是Student 也可以是people
	fmt.Println("t:", t)
}
