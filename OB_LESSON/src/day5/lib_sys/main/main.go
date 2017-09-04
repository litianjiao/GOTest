package main

import (
	"../model"
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
func SysInit() {

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
func Menu() {
	var input int

	fmt.Println()
	fmt.Println("1.借书功能")
	fmt.Println("2.学生信息管理")
	fmt.Println("3.书籍管理")

	fmt.Scanf("%d", &input)
	switch input {
	case FUNC_BOOK:
		model.BookMenu()
	case FUNC_STU:
		model.StuMenu()
	case FUNC_BOOKLIST:
		model.BookHandle()
	default:

	}
}

func main() {
	SysInit()
	for {
		Menu()
	}
}
