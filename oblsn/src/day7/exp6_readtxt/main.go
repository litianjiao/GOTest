package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	inputFile := "file1.txt"
	outputFile := "file1_cpy.txt"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		//Fprint函数指定了writer是终端标准输出还是网络或者别的
		fmt.Fprintf(os.Stderr, "File err:%s\n", err)
	}

	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0x664)
}
