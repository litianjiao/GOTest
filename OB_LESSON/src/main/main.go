package main

import "fmt"

/*
******************************************************************
  * @brief   get   factorial
  * @param
  * @ret
  * @author    TRoy
  * @date      2017/7/7 16:36
******************************************************************
*/
func sum(n int) uint64 {
	var s uint64 = 1
	var sum uint64 = 0
	for i := 1; i <= n; i++ {
		s = s * uint64(i)
		sum += s
	}
	return sum
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	s := sum(n)
	fmt.Printf("sum is %d", s)
}
