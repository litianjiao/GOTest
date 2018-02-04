package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int
	ch = make(chan int, 10)
	ch2 := make(chan int, 10)
	//每隔一秒分别向两个管道其中一个丢一个数据
	go func() {
		var i int
		for {
			ch <- i
			time.Sleep(time.Second)
			ch2 <- i * i
			time.Sleep(time.Second)
			i++
		}
	}()
	//取数据过程
	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case v := <-ch2:
			fmt.Println(v)
		case <-time.After(time.Second):
			fmt.Println("get data timeout")
			time.Sleep(time.Second)
		}
	}
}
