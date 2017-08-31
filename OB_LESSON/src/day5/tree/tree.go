package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	score float32
	left  *Student
	right *Student
}

func trans(node *Student) {
	if node == nil {
		return
	}
	fmt.Println(node)

	trans(node.left)
	trans(node.right)
}

func main() {
	var root *Student = new(Student)
	root.Name = "root"
	root.Age = 20
	root.score = 100

	var left1 *Student = new(Student)
	left1.Name = "node1"
	left1.Age = 24
	left1.score = 100

	root.left = left1

	var right1 *Student = new(Student)
	right1.Name = "node2"
	right1.Age = 24
	right1.score = 100

	root.right = right1

	var left2 *Student = new(Student)
	left2.Name = "node3"
	left2.Age = 24
	left2.score = 100

	left1.left = left2

	trans(root)
}
