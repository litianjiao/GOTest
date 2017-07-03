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

