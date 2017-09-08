package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string `json:"student_name"` //TAG的使用，用自己制定的别名打包
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func main() {
	var stu Student = Student{
		Name:  "stu01",
		Age:   18,
		Score: 100,
	}

	data, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("json encode stu failed err:", err)
		return
	}
	fmt.Println(string(data))
}
