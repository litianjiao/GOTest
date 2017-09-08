package work

import "fmt"

/*
******************************************************************
  * @brief   回文 英文解法
  * @param
  * @ret
  * @author    TRoy
  * @date      2017/7/21 9:35
******************************************************************

func process(str string) bool {

	for i := 0; i < len(str); i++ {
		if i == len(str) {
			break
		}
		last := len(str) - i - 1
		if str[i] != str[last] {
			return false
		}
	}
	return true
}
*/
/*
******************************************************************
  * @brief     回文 中文解法
  * @param
  * @ret
  * @author    TRoy
  * @date      2017/7/21 9:36
******************************************************************
*/
func process(str string) bool {
	t := []rune(str)
	len := len(t)

	for i, _ := range t {
		if i == len/2 {
			break
		}
		last := len - i - 1
		if t[i] != t[last] {
			return false
		}

	}
	return true
}

func main() {
	var str string
	fmt.Scanf("%sd", &str)
	if process(str) {
		fmt.Println("yes,it is ")
	} else {
		fmt.Println("no,it is not")
	}

}

//如上文range用法 i表示index 索引 省略的是c表示多个字节组成的一个rune类型字符，比如中文（3）
//不同的range有不同适用场景，如使用range来遍历字典的时候，返回键值对。
//kvs := map[string]string{"a": "apple", "b": "banana"}
//for k, v := range kvs {
//	fmt.Printf("%s -> %s\n", k, v)
//}
// range 用来遍历数组和切片的时候返回索引和元素值
// 如果我们不要关心索引可以使用一个下划线(_)来忽略这个返回值
// 当然我们有的时候也需要这个索引
//nums := []int{0,1,2,3,4}
//for i, num := range nums {
//	if num == 3 {
//		fmt.Println("index:", i)
//	}
//}
