package main

import (
	"fmt"
	"sort"
)

func map_sort() {
	var a map[int]int
	a = make(map[int]int, 5)

	a[8] = 10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[18] = 10

	var keys []int
	for k, v := range a {
		keys = append(keys, k)
		fmt.Println(k, v)
	}

	sort.Ints(keys)
	fmt.Println("after sort:")
	for _, v := range keys {
		fmt.Println(v, a[v])
	}
}

func map_reverse() {
	var a map[string]int
	var b map[int]string
	a = make(map[string]int, 5)
	b = make(map[int]string, 5)
	a["abc"] = 100
	a["des"] = 1
	fmt.Println(a)
	for k, v := range a {
		b[v] = k
	}
	fmt.Println("after reverse:")
	fmt.Println(b)
}

func main() {
	//map_sort()
	map_reverse()
}
