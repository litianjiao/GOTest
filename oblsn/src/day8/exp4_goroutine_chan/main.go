package main

import (
	"fmt"
	"time"
)
//管道容量为10 写10个进去便阻塞了
func write(ch chan int) {
	for i := 0; i < 100; i++ {
		ch<-i
		fmt.Println("put data:",i)
	}
}
//之后管道一个一个地读出来，
// write只能等管道读出来以后一个一个写进去

func read(ch chan int) {
	for{
		var b int
		b=<-ch
		fmt.Println(b)
		time.Sleep(time.Second)
	}
}



func main() {
	intChan:=make(chan int,10)
	go write(intChan)
	go read(intChan)

	time.Sleep(10*time.Second)
}
