package model

import (
	"errors"
	"fmt"
	"time"
)

var ErrBookNotFound = errors.New("Book Not Found")

type Student struct {
	Name  string
	Grade string
	ID    string
	Sex   string
	books []*BorrowItem
	//BookList map[string]*BookRecord
}

type BookRecord struct {
	start time.Time
	end   time.Time
	days  int
	book  *Book
}
type BorrowItem struct {
	book *Book
	num  int
}

/*
******************************************************************
  * @brief  学生信息管理界面
  * @param
  * @ret
  * @author    Troy
  * @date      2017/9/5 1:46
******************************************************************
*/
func StuMenu() {
	fmt.Println()
	fmt.Println("1.录入学生")
	fmt.Println("2.学生查询")
	fmt.Println("3.返回")

	var (
		t int
	)

	fmt.Scanf("%d", &t)
	switch t {
	case 1:
		stu := Student{
			Name:  "a",
			Grade: "1",
			ID:    "111",
			Sex:   "man",
		}
		fmt.Println("请输入姓名")
		fmt.Scanln(&stu.Name)
		fmt.Println("请输入年级")
		fmt.Scanln(&stu.Grade)
		fmt.Println("请输入学号")
		fmt.Scanln(&stu.ID)
		fmt.Println("请输入性别")
		fmt.Scanln(&stu.Sex)
		CreateStu(stu.Name, stu.Grade, stu.ID, stu.Sex)

		fmt.Printf("创建完成")
	}
}

/*
******************************************************************
  * @brief  录入学生
  * @param
  * @ret
  * @author    Troy
  * @date      2017/9/5 1:46
******************************************************************
*/
func CreateStu(name, grade, id, sex string) *Student {
	stu := &Student{
		Name:  name,
		Grade: grade,
		ID:    id,
		Sex:   sex,
	}
	return stu
}

/*
******************************************************************
  * @brief  获取学生个人信息
  * @param
  * @ret
  * @author    Troy
  * @date      2017/9/5 1:44
******************************************************************
*/
func (p Student) GetStu() {
	fmt.Println("学生信息:")
	fmt.Printf("学生姓名：%p\n", p.Name)
	fmt.Printf("学生年级：%p\n", p.Grade)
	fmt.Printf("学生ID：%p\n", p.ID)
	fmt.Printf("学生性别：%p\n", p.Sex)
}

/*
******************************************************************
  * @brief  修改学生个人信息
  * @param
  * @ret
  * @author    Troy
  * @date      2017/9/5 1:45
******************************************************************
*/
func (s *Student) EditStu() (err error) {
	return
}

func (s *Student) AddBook(b *BorrowItem) {
	s.books = append(s.books, b)
}
func (s *Student) GetBookList() []*BorrowItem {
	return s.books
}

func (s *Student) DelBook(b *BorrowItem) (err error) {

	for i := 0; i < len(s.books); i++ {
		if s.books[i].book.Name == b.book.Name {
			if b.num == s.books[i].num {
				front := s.books[0:1]
				left := s.books[i+1:]
				front = append(front, left...)
				s.books = front
				return
			}
			s.books[i].num -= b.num
		}
	}
	err = ErrBookNotFound
	return
}
