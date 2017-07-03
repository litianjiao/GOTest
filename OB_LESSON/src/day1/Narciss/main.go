package Narciss

import (
	"fmt"
)

/*
******************************************************************
  * @brief     求101~200的所有质数
  * @param
  * @ret
  * @author
  * @date
******************************************************************
*/
func main() {

	cnt_pri := 0
	for i := 101; i < 200; i++ {
		pri := 0
		for j := 2; j < i; j++ {
			if i%j == 0 {
				pri++
			}
		}
		if pri == 0 {
			cnt_pri++
			fmt.Printf("%d is a pri\n", i)
		}

	}

	fmt.Printf("the num of pri is %d", cnt_pri)
}

/*
******************************************************************
  * @brief     求100~1000的水仙花数
  * @param
  * @ret
  * @author
  * @date
******************************************************************

var (
	m   int
	j   int
	i   int
	s   int
	pri int
)
func main() {

	for n := 100; n < 1000; n++ {
		m = n
		j = 0
		for m != 0 {
			i = m % 10
			j += i * i * i
			m /= 10
		}
		if j == n {
			fmt.Printf("the %d is a Narcissistic\n", n)

		}
	}

}
*/
