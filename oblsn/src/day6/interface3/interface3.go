package main

import (
	"fmt"
)

type Student struct {
	Name string
	Sex  string
}

func convert(a interface{}) {
	b, ok := a.(Student)
	if !ok {
		fmt.Println("convert failed")
	}
	fmt.Println(b)
}

// 类型断言，将对象的原类型判断打印出来
//用可变参数 切片类型 来获取多个参数
func Classfiy(items ...interface{}) {
	for index, v := range items {
		switch v.(type) {
		case bool:
			fmt.Printf("%d params is bool,val is %v\n", index, v)
		case int, int64, int32:
			fmt.Printf("%d params is int,val is %v\n", index, v)
		case float32, float64:
			fmt.Printf("%d params is float,val is %v\n", index, v)
		case string:
			fmt.Printf("%d params is string,val is %v\n", index, v)
		case Student:
			fmt.Printf("%d params is Student,val is %v\n", index, v)
		case *Student:
			fmt.Printf("%d params is *Student,val is %v\n", index, v)
		}
	}
}

func main() {
	var a Student = Student{
		Name: "xiaoming",
		Sex:  "male",
	}
	convert(a)
	Classfiy(100, 2.2, "33333", a, &a)
}
