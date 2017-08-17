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
)

func find(p_str *string) (lost uint8, err bool) {
	var (
		sum uint8
		i   int
	)
	err = false
	str := &p_str
	str1 := []byte(**str)
	len := len(str1) - 1
	if len > 14 || len < 13 {
		err = true
		return
	}
	if len == 13 {
		for i = 0; i < len; i++ {
			str1[i] -= 48
			sum += uint8(str1[i])
		}
		lost = 55 - sum + 9
	} else if len == 14 {
		for i = 0; i < len; i++ {
			str1[i] -= 48
			sum += str1[i]
		}
		lost = 55 - sum
	}

	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		log.Println("输入流转换异常", err)
		os.Exit(-2)
	}
	lost, f_err := find(&str)
	if f_err {
		log.Println("输入有误")
		os.Exit(-1)
	}
	fmt.Printf("lost boy is %d", lost)
}
