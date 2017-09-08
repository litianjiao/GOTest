package main

import "fmt"
import "strings"
//闭包的例子
//闭包的第一个参数是public para 通过初始化时变化
// 第二个参数是实际受用对象
func makeSuffixFunc(suffix string)func(string) string {
	return func(name string)string{
		if !strings.HasSuffix(name,suffix){
			return name+suffix
		}
		return name
	}
}

func main()  {
	func1:=makeSuffixFunc(".avi")
	func2:=makeSuffixFunc(".mp4")
	fmt.Println(func1("IPZ-929"))
	fmt.Println(func2("LXVS-036"))
}