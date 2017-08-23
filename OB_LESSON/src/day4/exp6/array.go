package main

import "fmt"

// slice typedef
type slice struct {
	ptr *[50]int
	len int
	cap int
}

// make slice
func make_slice(s slice, cap int) slice {
	s.cap = cap
	s.len = 0
	s.ptr = new([50]int)
	return s
}

// modify slice
func modify(s slice) {
	s.ptr[1] = 100
}

func range_multi_array() {
	var d [2][3]int = [...][3]int{{1, 2, 3}, {4, 5, 6}}

	for k1, v1 := range d {
		for k2, v2 := range v1 {
			fmt.Printf("(%d,%d)=%d\n", k1, k2, v2)
		}
		fmt.Println()
	}
}


//
func test() {
	var s1 slice
	s1 = make_slice(s1, 10)
	s1.ptr[0] = 20
	modify(s1)
	fmt.Println(s1.ptr)
}

func main() {

	test()
}
