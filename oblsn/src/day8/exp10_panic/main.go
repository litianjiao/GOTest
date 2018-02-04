package main

import (
	"fmt"
	"runtime"
	"time"
)

//如何捕获panic

//出问题的打panic
func test()  {

	defer func() {

		//如果err！=空，即表示出问题了
		if err:=recover();err!=nil{
			fmt.Println("panic:",err)
		}
	}()
	//此处定义了map未初始化
    var m map[string]int
	m["num"]=100 //IDE其实已经指出问题了
	}
//正常工作的协程不受影响 继续干活
func run1()  {
		for{
			fmt.Println("im run")
		    time.Sleep(time.Second)
			}
}
func main()  {
	num:=runtime.NumCPU()
	go run1()//我不受影响 继续干活
	runtime.GOMAXPROCS(num-1)
	for i := 0; i <2; i++ {
		go test() //我们两兄弟被panic了 但是打印了为啥

	}
	time.Sleep(time.Second*10000)
}
