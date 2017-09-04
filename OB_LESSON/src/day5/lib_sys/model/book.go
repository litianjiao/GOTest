package model

import (
	"errors"
	"fmt"

)


const (
	DEFAULT = iota
	BORROW
	BACK
	SEARCH
	UP
)

var (
	ErrStockNotEnough = errors.New("Stock is not enough")
)

type Book struct {
	Name    string
	Copy    int
	Author  string
	ISBN    string
	Date    string
	Publish string
	Flag    int
}

/*
******************************************************************
  * @brief  BookMenu
  * @param
  * @ret
  * @author    Troy
  * @date      2017/9/4 10:13
******************************************************************
*/
func BookMenu() {
	var user_type int
	fmt.Println()
	fmt.Println("1.借书")
	fmt.Println("2.还书")
	fmt.Println("3.书籍查询")
	fmt.Println("4.返回")
	fmt.Scanf("%d", &user_type)
	switch user_type {
	case BORROW:
		borrowBook()
	case BACK:
		backBook()
	case SEARCH:
		searchBook()
	case UP:
		return

	default:

	}

}

/*
******************************************************************
  * @brief    BookHandle
  * @param
  * @ret
  * @author    Troy
  * @date      2017/9/4 10:12
******************************************************************
*/
func (s *Student) BookHandle() {

	for {
		fmt.Println()
		fmt.Println("1.录入")
		fmt.Println("2.书籍查询")
		fmt.Println("3.书籍借还清单")
		fmt.Println("4.返回")
		var b_t int
		fmt.Scanf("%p", &b_t)
		switch b_t {
		case 1:
			CreateBook("TGPL01", 10, "Duonuowan", "2010")
		case 2:
			fmt.Println(s.GetBookList())
		case 3:
		case 4:
			return
		default:
		}
	}
}

func borrowBook() {

}

func backBook() {

}

func searchBook() {

}

func CreateBook(name string, copy int, author, createTime string) (b *Book) {
	b = &Book{
		Name:    name,
		Copy:    copy,
		Author:  author,
		Publish: createTime,
	}
	return
}

func (b *Book) canBorrow(c int) bool {
	return b.Copy >= c
}

func (b *Book) Borrow(c int) (err error) {
	if b.canBorrow(c) == false {
		err = ErrStockNotEnough
		return
	}
	b.Copy -= c
	return
}

func (b *Book) Back(c int) (err error) {

	return
}
