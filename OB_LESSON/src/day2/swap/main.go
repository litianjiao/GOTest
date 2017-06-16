package main

import (
	"fmt"
)

/*
func swap(a int, b int) {
	tmp := a
	a = b
	b = tmp
	return
}
这么做只是交换了swap函数中a b的副本 a1 b1的值，并未影响外面传进来参数的值
如果要改 ，需要操作地址所指向内存的值
*/

func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
	return
}
func main() {
	num1 := 100
	num2 := 200
	swap(&num1, &num2)
	fmt.Println("first=", num1)
	fmt.Println("second=", num2)
}
