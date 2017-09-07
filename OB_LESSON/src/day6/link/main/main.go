package main

import (
	"../linklist"
	"fmt"
)

func main() {

	var link1 linklist.Link

	for i := 0; i < 10; i++ {
		//link1.InsertHead(i)
		//link1.InsertTail(i)
		//link1.InsertHead(fmt.Sprintf("node %d", i))
		link1.InsertTail(fmt.Sprintf("node %d", i))
	}

	link1.Trans()
}
