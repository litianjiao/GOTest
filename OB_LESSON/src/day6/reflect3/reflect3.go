package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

//反射：
//反射的操作不能做静态类型检查，而且大量反射的代码通常难以理解。
//总是需要小心翼翼地为每个导出的类型和其它接受interface{}或reflect.Value类型参数的函数维护说明文档。
//第三个原因，基于反射的代码通常比正常的代码运行速度慢一到两个数量级。
//对于一个典型的项目，大部分函数的性能和程序的整体性能关系不大，所以使用反射可能会使程序更加清晰。
//测试是一个特别适合使用反射的场景，因为每个测试的数据集都很小。但是对于性能关键路径的函数，最好避免使用反射。

type Student struct {
	Name  string `json:"student_name"` //Tag
	Age   int
	Score float32
	sex   string
}

func (s Student) Print() {
	fmt.Println("----X---")
	fmt.Println(s)
	fmt.Println("----y---")
}

func (s Student) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.sex = sex
}
func (s Student) SetAge(age int) {
	s.Age = age
}

//通过反射 操作结构体
func TestStruct(a interface{}) {
	//tye := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Struct {
		fmt.Println("expect struct")
	}
	//是否是指针 && 这个指针所指向的对象是否是个结构体
	// if kd != reflect.Ptr && val.Elem().Kind() == reflect.Struct {
	// 	fmt.Println("expect struct")
	// 	return
	// }
	num := val.NumField()
	//	val.Field(0).SetString("xiaoxiaoming")
	for i := 0; i < num; i++ {
		fmt.Printf("%d %v\n", i, val.Field(i).Kind())
	}
	fmt.Printf("struct has %d fields\n", num) ///结构体字段数
	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod) ///结构体定义方法数
	var params []reflect.Value

	val.Method(2).Call(params)
}

//通过反射修改机构体
func TestStruct2(a interface{}) {
	tye := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Ptr && val.Elem().Kind() == reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	num := val.Elem().NumField()
	val.Elem().Field(0).SetString("stu1000")
	for i := 0; i < num; i++ {
		fmt.Printf("%d %v\n", i, val.Elem().Field(i).Kind())
	}

	fmt.Printf("struct has %d fields\n", num)
	//TAG是type的方法
	tag := tye.Elem().Field(0).Tag.Get("json")
	fmt.Printf("tag=%s\n", tag)

	numOfMethod := val.Elem().NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)
	var params []reflect.Value
	val.Elem().Method(0).Call(params)
}

func main() {
	var a Student = Student{
		Name:  "xiaoming",
		Age:   20,
		Score: 95.1,
	}
	result, _ := json.Marshal(a)
	fmt.Println("json reslut:", string(result))
	//TestStruct(a)
	TestStruct2(&a)
}
