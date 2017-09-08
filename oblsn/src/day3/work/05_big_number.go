package work

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
******************************************************************
  * @brief     大数相加
  * @param
  * @ret
  * @author    TRoy
  * @date      2017/7/22 7:53
******************************************************************
*/
/*两个大数先转化为字符串再通过for循环从index最大的地方一位一位相加*/
func add(str1, str2 string) (result string) {
	/*如果俩字符串长度都为0 说明原字符串不存在*/
	if len(str1) == 0 && len(str2) == 0 {
		result = "0"
		return
	}
	/*位数较少的，高位补0*/
	if len(str1) < len(str2) {
		str1 = strings.Repeat("0", len(str2)-len(str1)) + str1
	} else if len(str1) > len(str2) {
		str2 = strings.Repeat("0", len(str1)-len(str2)) + str2
	}
	var index = len(str1) - 1 //index最大的地方其实相当于个位，记得要减一
	var left int //进位
	/*从后往前加*/


	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	res, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("err is:", err)
	}
	/*拆分*/
	strslice := strings.Split(string(res), "+")
	if len(strslice) != 2 {
		fmt.Println("please input a+b")
		return
	}
	//很重要的一步就是俩化成字符串的数不能以0000123这样前面有0的形式
	strNumber1 := strings.TrimSpace(strslice[0])
	strNumber2 := strings.TrimSpace(strslice[1])
	fmt.Println(add(strNumber1, strNumber2))

}
