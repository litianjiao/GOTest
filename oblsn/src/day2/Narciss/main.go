package Narciss

import (
	"fmt"
)

/*
******************************************************************
  * @brief     求100~1000的水仙花数
  * @param
  * @ret
  * @author
  * @date
******************************************************************/

/*
******************************************************************
  * @brief  isNarciss
  * @param   n int the num to judge
  * @ret     yesy or no
  * @author    TRoy
  * @date      2017/7/7 15:39
******************************************************************
*/
func isNarciss(n int) bool {
	var m, j, i int
	m = n
	for m != 0 {
		i = m % 10
		j += i * i * i
		m /= 10
	}
	if j == n {
		return true
	} else {
		return false
	}

}

func main() {
	var m, n int
	fmt.Scanf("%d,%d", &m, &n)

	for i := m; i < n; i++ {
		if isNarciss(i) == true {
			fmt.Printf("the %d is a Narcissistic\n", i)
		}
	}

}
