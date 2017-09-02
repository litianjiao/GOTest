package main

import (
	"fmt"
)

//Anonymous field 匿名字段

type Cart1 struct {
	name string
	age  int
}

type Cart2 struct {
	name string
	age  int
}
type Train struct {
	Cart1
	Cart2
}

func main() {
	var t Train

	t.Cart1.name = "train"
	//当匿名字段内字段不同时 可以简写
	//如t.age
	t.Cart1.age = 100
	fmt.Println(t)
}
