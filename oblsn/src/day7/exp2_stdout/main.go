package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("./nice.log", os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		fmt.Println("open file err \n")
		return
	}
	fmt.Fprintf(file, "do balance err\n")
	file.Close()
}
