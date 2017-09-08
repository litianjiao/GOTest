package main

import "fmt"
//导致panic异常的函数不会继续运行，但能正常返回。
//被延迟执行的匿名函数可以修改函数返回给调用者的返回值。
//所以在panic 函数执行后，函数准备返回result 这个变量，之后执行defer 的func，在这个func 里改变了result 的值
func main() {
    a := returnN()
    fmt.Println(a)
}

func returnN() (result int) {
    defer func() {
        if p := recover(); p != nil {
            result = p.(int)
        }
    }()
    panic(3)
}