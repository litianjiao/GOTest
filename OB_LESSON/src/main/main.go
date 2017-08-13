package main

import (
	"fmt"
	"smog"
)

func main() {
	fmt.Printf("please enter your SN.\n\r ")
	fmt.Println("yes")
	Device := smog.Loader()
	smog.Test_Process(&Device)

}

//fmt.Printf("create %s succ\ntype: %s,\nsn: %s,\n", Test_device.Name, Test_device.Type, Test_device.SN)
