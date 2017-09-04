package main

import (
	_ "../model"
	"fmt"
)

const (
	FUNC_BOOK     = 1
	FUNC_STU      = 2
	FUNC_BOOKLIST = 3
)

/*
******************************************************************
  * @brief      Init Sys
  * @param
  * @ret
  * @author    Troy
  * @date      2017/9/4 8:42
******************************************************************
*/
func SysInit() (stu *Student) {

	CreateStu("Xiaoming", "一年级", "1311020111", "男")
	return
}

/*
******************************************************************
  * @brief  main menu
  * @param
  * @ret
  * @author    Troy
  * @date      2017/9/4 8:59
******************************************************************
*/
func (p *User) Menu() {
	var input int
	fmt.Println()
	fmt.Println("1.借书功能")
	fmt.Println("2.学生信息管理")
	fmt.Println("3.书籍管理")

	fmt.Scanf("%d", &input)
	switch input {
	case FUNC_BOOK:
		User.BookMenu()
	case FUNC_STU:
		User.StuMenu()
	case FUNC_BOOKLIST:
		User.BookHandle()
	default:

	}
}

func main() {
	stu := SysInit()
	for {
		stu.Menu()
	}
}
