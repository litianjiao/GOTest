package main

import (
	"fmt"
	"math/rand"
)

type Student struct {
	Name  string
	Age   int
	score float32
	next  *Student
}

func trans(p *Student) {
	for p != nil {
		fmt.Println(*p)
		p = p.next
	}
	fmt.Println()
}
func insertTail(p *Student) {
	var tail = p
	for i := 0; i < 10; i++ {
		stu := Student{
			Name:  fmt.Sprintf("stu%d", i),
			Age:   rand.Intn(100),
			score: rand.Float32() * 100,
		}
		tail.next = &stu
		tail = &stu
	}
}
func insertHead(p **Student) {

}

func main() {
	var head *Student = new(Student)
	head.Name = "Xiaoming"

	head.Age = 10
	head.score = 99

	insertTail(&head)
	trans(&head)

}
