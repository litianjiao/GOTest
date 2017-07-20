package exam

import "fmt"

/*
******************************************************************
  * @brief   break &continue
  * @param
  * @ret
  * @author    TRoy
  * @date      2017/7/19 23:36
******************************************************************
*/
/*GOlang中string类型的编码都是utf-8，可以看到一个英文字符是1字节，一个中文字是3字节*/
func main() {
	str := "he234 world, 芭娜娜"
	for index, val := range str {
		//如果此处有continue的话 那就是部分打印了，当index满足某条件进入continue时，直接到下一次for continue以下的print函数并不执行
		if index < 2 {
			continue //  equa  if index<2 go to =>index++ =>range str[3]
		}
		if index > 5 {
			break
		}
		//以上的方式制约了打印数据的开头和结尾
		fmt.Printf("index[%d] val[%c] len[%d]\n", index, val, len([]byte(string(val))))
	}

}
