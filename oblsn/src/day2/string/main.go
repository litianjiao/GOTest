package main

import (
	"fmt"
)

/*
******************************************************************
  * @brief     字符串顺序
  * @param
  * @ret
  * @author
  * @date
******************************************************************
*/
func reverse(str string) string {
	var result string
	strlen := len(str)
	for i := 0; i < strlen; i++ {
		result = result + fmt.Sprintf("%c", str[strlen-1-i])
	}
	return result
}

func main() {
	var (
		str1 = "hello"
		str2 = "world"
	)
	str3 := fmt.Sprintf("%s %s\n", str1, str2)
	n := len(str3)
	fmt.Printf(str3)
	fmt.Printf("len(str)=%d\n", n)

	substr := str3[0:7]
	fmt.Println(substr)

	result := reverse(str3)
	fmt.Println(result)
}
