package main

import (
	"time"
	"fmt"
)

func main() {
	t:=time.NewTicker(5*time.Second)
	for  v := range t.C {
		fmt.Println("time afer 5 second,now is",v)
	}
}
