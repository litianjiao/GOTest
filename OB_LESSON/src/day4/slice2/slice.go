package main

import "fmt"

/*
******************************************************************
  * @brief    reverse reverses a slice of ints in place.
  * @param  s []int
  * @ret
  * @author    Troy
  * @date      2017/8/23 22:49
          j=len(s)-1 for j as index of slice.(下标作用)
******************************************************************
*/
func s_reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

/*
******************************************************************
  * @brief     s_equal
  * @param  x,y []string
  * @ret    bool
  * @author    Troy
  * @date      2017/8/23 22:58
******************************************************************
*/
func s_equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6} //array
	s := []int{0, 1, 2, 3, 4, 5}       //slice of array a
	s_reverse(a[:])                    //use slice for modify source array data
	s_reverse(s[:2])
	fmt.Println(s)
	s_reverse(s[2:])
	fmt.Println(s)
	s_reverse(s)
	fmt.Println(s)
	fmt.Println(a)
}
