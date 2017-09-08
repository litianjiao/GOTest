package main

import (
	"fmt"
	"math/rand"
)

type Student struct {
	Name  string
	Age   int
	Score float32
	next  *Student
}

//遍历
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
			Score: rand.Float32() * 100,
		}
		tail.next = &stu
		tail = &stu
	}
}

//改变变量的值，传指针进func
//改变指针变量的值，就传这个指针变量的指针进func
//因为每次操作都要将
func insertHead(p **Student) {
	for i := 0; i < 10; i++ {
		stu := Student{
			Name:  fmt.Sprintf("stu%d", i),
			Age:   rand.Intn(100),
			Score: rand.Float32() * 100,
		}
		stu.next = *p
		*p = &stu
	}
}

func delNode(p *Student) {
	var prev *Student = p
	for p != nil {
		if p.Name == "stu4" {
			prev.next = p.next
			break
		}
		prev = p
		p = p.next
	}
}

func addNode(p *Student, newNode *Student) {
	for p != nil {
		if p.Name == "stu9" {
			newNode.next = p.next
			p.next = newNode
			break
		}
		p = p.next
	}
}

func main() {
	var head *Student = new(Student)

	head.Name = "Xiaoming"
	head.Age = 10
	head.Score = 99
	fmt.Printf("head adr is %p\n", &head)
	//insertTail(head) //  尾部插入（xiaoming->stu0->stu1->stu2...）
	insertHead(&head) //  头部插入 (...stu2<-stu1<-stu0<-xiaoming)
	fmt.Println("before del")
	trans(head)
	delNode(head)
	fmt.Println("after del")
	trans(head)

}
