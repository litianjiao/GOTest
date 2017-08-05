package work

import "fmt"

/*
******************************************************************
  * @brief      完数
  * @param
  * @ret
  * @author    TRoy
  * @date      2017/7/20 22:47
******************************************************************
*/
func perfect(n int) bool {
	var sum int = 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return n == sum
}
func process(n int) {
	for i := 1; i < n+1; i++ {
		if perfect(i) {
			fmt.Println(i)
		}
	}
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	process(n)
}
