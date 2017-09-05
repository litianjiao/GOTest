/*
    幼儿园有10个小朋友，把1 – 20中所有奇数数字的卡片分别发给这10个小朋友。
    集合的时候小朋友们排成一排，组成一个由数字组成的字符串，如59731315….11917，
	但是老师发现少了一个小朋友，你能帮忙找出少掉的那个小朋友吗？
*/

//思路：第一步 判断是少了一位的还是少了两位的小朋友 那字符串长度就对应的是13或者14
// 1 3 5 7 9  11 13 15 17 19
//我们已转化为整数的值总和是一定的（55），因此根据少的数目补位
// 2.或者接着统计13579分别出现的次数
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
******************************************************************
  * @brief     find lost children
  * @param
  * @ret
  * @author
  * @date      2017/8/17 21:40
******************************************************************
*/
func find(p_str *string) (str_lost string, err bool) {
	var (
		sum  uint8
		lost uint8
	)
	err = false
	str := &p_str
	str1 := []byte(**str)
	len := len(str1)
	fmt.Printf("str len is %d\n", len)
	if len > 14 || len < 13 {
		err = true
		return
	}
	//去除杂质后遍历数组
	for _, v := range str1 {
		v -= 48
		sum += uint8(v)
	}
	if len == 13 {
		lost = 55 - sum + 9
		fmt.Printf("001 %d\n", lost)

	} else if len == 14 {

		lost = 55 - sum
		fmt.Printf("002 %d\n", lost)
	}
	str_lost = strconv.Itoa(int(lost))
	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		log.Println("输入流转换异常", err)
		os.Exit(-2)
	}
	str = strings.Replace(str, " ", "", -1)  //去除空格
	str = strings.Replace(str, "\n", "", -1) //去除\n
	str = strings.Replace(str, "\r", "", -1) //去除\r
	lost, f_err := find(&str)
	if f_err {
		log.Println("输入有误")
		os.Exit(-1)
	}
	fmt.Println("lost boy is ", lost)
}
