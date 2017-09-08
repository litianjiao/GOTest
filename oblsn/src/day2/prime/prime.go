package prime

import (
	"fmt"
	"math"
)

/*
******************************************************************
  * @brief     isPrime
  * @param     n int
  * @ret       ture false
  * @author    TRoy
  * @date      2017/7/3 23:14
******************************************************************
*/
func isPrime(n int) bool {
	for i := 2; i <= int(math.Sqrt(n)); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

/*
******************************************************************
  * @brief    求101~200的所有质数
  * @param
  * @ret
  * @author    TRoy
  * @date      2017/7/3 23:10
******************************************************************
*/
func main() {
	var (
		n int
		m int
	)
	fmt.Scanf("%d%d", &n, &m)

	for i := n; i < m; i++ {
		if isPrime(i) == true {
			fmt.Printf("%d is a Prime\n", i)
			continue
		}
	}
}

//func main() {
//
//	cnt_pri := 0
//	for i := 101; i < 200; i++ {
//		pri := 0
//		for j := 2; j < i; j++ {
//			if i%j == 0 {
//				pri++
//			}
//		}
//		if pri == 0 {
//			cnt_pri++
//			fmt.Printf("%d is a pri\n", i)
//		}
//
//	}
//
//	fmt.Printf("the num of pri is %d", cnt_pri)
//}
