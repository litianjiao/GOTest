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
func modify1(a []int) {
	a[0] = 1000
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

func slice_test1(){
	var a = [10]int{1, 2, 3, 4}
	
		b := a[1:5]
		fmt.Printf("%p\n", b)
		fmt.Printf("%p\n", &a[1])
}

func slice_test2() {
	var b []int = []int{1, 2, 3, 4}
	modify1(b)
	fmt.Println(b)
}

func slice_test3()  {
	var a [5]int = [...]int{1, 2, 3, 4, 5}
	s := a[1:]
	fmt.Printf("before len[%d] cap[%d]\n", len(s), cap(s))
	s[1] = 100
	fmt.Printf("s=%p a[1]=%p\n", s, &a[1])
	fmt.Println("before a:", a)
	s=append(s,10,10,10)	
	fmt.Printf("after len[%d] cap[%d]\n", len(s), cap(s))
	//此时len变化，但是并没有超出cap，所以cap不变 此时还是4 作为数组a索引的slice地址不变
	s = append(s, 20,20,20)
	s[1]=300
	fmt.Println("after a:", a)
	fmt.Println(s)
	fmt.Printf("s=%p a[1]=%p\n", s, &a[1])
	//扩容之后一看到作为数组a索引的slice的首地址居然发生了变化，说明append函数在使用过程中，若cap变化，则地址会改变！
	//扩容时相当于cap*2
}

func testCopy() {
	
		var a []int = []int{1, 2, 3, 4, 5, 6}
		b := make([]int, 1)
	
		copy(b, a)
	
		fmt.Println(b)
}

func testString() {
		s := "hello world"
		s1 := s[0:5]
		s2 := s[6:]
	
		fmt.Println(s1)
		fmt.Println(s2)
	
}
func main() {

	//test()
	// slice_test3()
	// slice_test1()
	testString()
}
