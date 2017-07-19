package strings

import (
	"fmt"
	"strings"
)

//package main
//
//import (
//"fmt"
//"strings"
//)
//
//func main() {
//	str := "hello world hello world"
//	result := strings.Replace(str, "hello", "HELLO", 0)
//	fmt.Printf("%s", result)
//
//}

//func urlprocess(url string) string {
//	result := strings.HasPrefix(url, "http://")
//	if !result {
//		url = fmt.Sprintf("http://%s", url)
//	}
//	return url
//}
//
//func pathprocess(path string) string {
//	result := strings.HasPrefix(path, "/")
//	if !result {
//		path = fmt.Sprintf("/%s", path)
//	}
//	return path
//}
//func main() {
//	var (
//		url  string
//		path string
//	)
//	fmt.Scanf("%s%s", &url, &path)
//	url = urlprocess(url)
//	path = pathprocess(path)
//	println(url)
//	println(path)
//
//}
/*
******************************************************************
  * @brief     strfix
  * @param     (url string)
  * @ret        string
  * @author    TRoy
  * @date      2017/7/19 20:46
******************************************************************
*/
func strfix(s string, suff string) string {
	res := strings.HasSuffix(s, suff)
	if !res {
		s += suff
	}
	return s
}
func main() {
	var (
		s1 string
	//	s2 string
	)
	s1 = "hello world aacc \n\t\r\n\n"
	//fmt.Scanf("%s%s", &s1, &s2)
	//s1 = strfix(s1, s2)
	res := strings.Replace(s1, "hello", "Hello", 1)
	fmt.Println("replace:", res)
	cnt := strings.Count(s1, "l")
	fmt.Println("count:", cnt)
	res = strings.TrimSpace(s1)
	fmt.Println("trimspace:", res)
}
