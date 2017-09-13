package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCnt struct {
	NumCnt   int
	SpaceCnt int
	CharCnt  int
	OtherCnt int
}

func main() {
	file, err := os.Open("./nice.log")
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer file.Close()
	var cnt CharCnt
	reader := bufio.NewReader(file)
	for {

		str, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("read over.")
			break
		}
		if err != nil {
			fmt.Println("read string failed,err:", err)
			break
		}
		runeArr := []rune(str)
		for _, v := range runeArr {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				cnt.CharCnt++
			case v >= '0' && v <= '9':
				cnt.NumCnt++
			case v == ' ' || v == '\t':
				cnt.SpaceCnt++
			default:
				cnt.OtherCnt++

			}
		}
		//fmt.Printf("read str succ,ret: %s \n", str)
	}
	fmt.Printf("char cnt:%d\n", cnt.CharCnt)
	fmt.Printf("num cnt:%d\n", cnt.NumCnt)
	fmt.Printf("space cnt:%d\n", cnt.SpaceCnt)
	fmt.Printf("other cnt:%d\n", cnt.OtherCnt)

}
