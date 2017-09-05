package main

import (
	"fmt"
	"smog"
)

func main() {
	fmt.Printf("please enter your SN.\n\r ")
	Device := smog.Loader()
	smog.Test_Process(&Device)

}
