package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string
	Age   int
	Score float32
}

func TestStruct(a interface{}) {
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
}

//复习for range 这里面的v只是值拷贝 如果要改变原变量内容 要么原数组类型变成slice
//要么把v改成astr[i] 因为astr[i]是指针 是引用类型
func Test2() {
	aStr := [3]int{0, 1, 2}
	for i, v := range aStr {
		if i == 0 {
			aStr[1], aStr[2] = 99, 999
			fmt.Printf("befor astr[%d]=%d\n", i, v)
		}
		aStr[i] = aStr[i] + 100
		//aStr[i] = 100+v
		fmt.Printf("after astr[%d]=%d\n", i, aStr[i])
	}
	fmt.Println(aStr)
}
func main() {
	var a Student = Student{
		Name:  "xiaoming",
		Age:   20,
		Score: 95.1,
	}

	TestStruct(a)
	fmt.Println(a)
	Test2()
}
