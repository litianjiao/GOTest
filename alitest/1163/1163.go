package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	n      int
	maxSum *int
)

//将控制台的输入(每行的数据)转换成整形数组
func convsStrToInt(num string) (array []int, err error) {

	//去掉换行符后以逗号分割字符串
	arrayTem := strings.Split(strings.TrimSpace(num), ",")

	//判断输入个数的合法性
	if len(arrayTem) != n {
		err = errors.New("输入个数超过矩阵列数或检查是否用逗号隔开")
		return
	}
	return
}

func main() {

	fmt.Scanln(&n)
	if n == 0 {
		log.Println("参数非法")
		return
	}
	num_input := make([][]int, 0, 0)
	for {
		inputReader := bufio.NewReader(os.Stdin)
		input, err := inputReader.ReadString('\n')
		if err != nil {
			log.Println("err in inputs")
			break
		}
		arrInt, err_conv := convsStrToInt(input)
		if err_conv != nil {
			log.Println(err_conv)
			return
		}
		num_input = append(num_input, arrInt)
	}

	//maxSum=D[n]
	for _, v := range num_input[n-1] {
		fmt.Println(v)
	}
}
