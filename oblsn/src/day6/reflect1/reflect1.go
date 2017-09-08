package main

import "reflect"
import "fmt"

type Student struct {
	Name string
	Age  int
}

func test(b interface{}) {
	t := reflect.TypeOf(b) //type of  关联变量运行环境
	fmt.Println(t)

	v := reflect.ValueOf(b)
	k := v.Kind() //kind of 不关联变量运行环境 ，根据所定义的来
	fmt.Println(k)

	iv := v.Interface()
	stu, ok := iv.(Student)
	if ok {
		fmt.Printf("%v %T\n", stu, stu)
	}
}

//只读 不修改原先的变量
func testInt(b interface{}) {
	val := reflect.ValueOf(b)
	c := val.Int() //如果这个潜在值val原类型不是val则会panic
	//因为是值类型，所以传进来的是副本 所以无法改变原值
	fmt.Printf("get val interface %d\n", c)
}

//如果要修改原先的变量，则需使用指针方式  Elem相当于指针方式
func testpInt(b interface{}) {
	val := reflect.ValueOf(b)

	c := val.Elem().Int()
	fmt.Printf("get val interface{} %d\n", c)
	fmt.Printf("Int val:%d\n", val.Elem().Int())

}

func main() {
	var a Student = Student{
		Name: "xiaoming",
		Age:  20,
	}
	test(a)
	testInt(1111)
	var b string = "1"
	//b = 200
	testpInt(&b)
	fmt.Println(b)
}
