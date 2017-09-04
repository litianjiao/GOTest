package model

import (
	"time"
	"fmt"
)

type Student struct {
	Name     string
	Grade    string
	ID       string
	Sex      string
	BookList map[string]*BookRecord
}

type BookRecord struct {
	start time.Time
	end   time.Time
	days  int
	book  *Book
}

func StuMenu() {
	fmt.Println()
	fmt.Println("1.录入学生")
	fmt.Println("2.学生查询")
	fmt.Println("3.返回")

	var t int
	fmt.Scanf("%d",&t)
	switch t {
	case 1:
		fmt.Println("请输入姓名，年级，学号，性别")
		
		CreateStu("")
	}
}

func CreateStu(name, grade, id, sex string) *Student {
	stu := &Student{
		Name:  name,
		Grade: grade,
		ID:    id,
		Sex:   sex,
	}
	return stu
}
