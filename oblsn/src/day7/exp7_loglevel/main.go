package main

import (
	"flag"
	"fmt"
)

func main() {
	var filePath string
	var logLevel int
	//adr of store the val ..name ..default val  ..  point
	flag.StringVar(&filePath, "c", "", "please input path")
	flag.IntVar(&logLevel, "d", 10, "please input log level")
	//after all flags are defined and before flags are accessed by the program.
	flag.Parse()

	fmt.Println("path:", filePath)
	fmt.Println("log level:", logLevel)
}

//USE -c -d for edit file path and log level
