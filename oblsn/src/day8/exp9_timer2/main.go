package main

import (
	"runtime"
	"time"
	"fmt"
)

func main() {
	num:=runtime.NumCPU()
	runtime.GOMAXPROCS(num-1)
	for i := 0; i < 1024; i++ {
		go func() {
			for{
				t:=time.NewTicker(time.Second)
				select {
				case<-t.C:
					fmt.Println("timeout")
				}
				t.Stop() //用完定时器以后一定要关，防止内存泄漏
			}
		}()
	}
	t
}
