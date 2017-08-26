package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Comparable interface {
	// Len is the number of elements in the collection.
	len() int

	// Less reports whether the element with
	// index i should sort before the element with index j.
	less(i, j int) bool

	// Swap swaps the elements with indexes i and j.
	swap(i, j int)

	show(i int) int
}
type IntSlice []int

func (p IntSlice) sort()              { sort(p) }
func (p IntSlice) len() int           { return len(p) }
func (p IntSlice) less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p IntSlice) show(i int) int     { return p[i] }

//检查是否有序
func isSorted(data Comparable) bool {
	n := data.len()
	for i := n - 1; i > 0; i-- {
		if data.less(i, i-1) {
			return false
		}
	}
	return true
}
func RandArray(n int) []int {
	arr := make([]int, n)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i <= n-1; i++ {
		arr[i] = rand.Intn(n)
	}
	return arr
}

//用引用类型【切片】来操作对应数组
//冒泡
//相邻元素比较【互换】
func bubbleSort(a []int) {
	var swap int
	for i := 0; i < len(a); i++ {
		for j := 1; j < len(a)-i; j++ {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
				swap++
			}
		}
	}
	fmt.Printf("has swap:%d\n", swap)
}

// func bubbleSort(a []int) {
// 	var swap int
// 	for i := 0; i < len(a); i++ {
// 		for j := 1; j < len(a)-i; j++ {
// 			if a[j] < a[j-1] {
// 				a[j], a[j-1] = a[j-1], a[j]
// 				swap++
// 			}
// 		}
// 	}
// 	fmt.Printf("has swap:%d\n", swap)
// }
//选择
//从第一个数开始 和其后面的无序数组进行比较，若最大（小）,将其放在第一个位置【变下标】
func selectSort(a []int) {
	var change int
	for i := 0; i < len(a); i++ {
		var min int = i
		for j := i + 1; j < len(a); j++ {
			if a[min] > a[j] {
				min = j
				change++
			}
		}
		a[i], a[min] = a[min], a[i]
	}
	fmt.Printf("has change:%d\n", change)
}

//插入
//从第二个数开始和其前面的有序数列进行比较，选择插入在哪
//根据确定位置的方法的不同 可衍生多种方法
func insertSort(a []int) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			//在这里用的是直接和前一个数冒泡
			if a[j] > a[j-1] {
				//和前一个数有序，则不换位置
				break
			}
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}

//快排
func quickSort(a Comparable, lo int, hi int) {
	if lo < hi {
		i, j := lo, hi
		pos := a.show((lo + hi) / 2)
		for i <= j {
			for a.show(i) < pos {
				i++
			}
			for a.show(j) > pos {
				j--
			}
			if i <= j {
				a.swap(i, j)
				i++
				j--
			}
		}
		if lo < j {
			quickSort(a, lo, j)
		}
		if hi > i {
			quickSort(a, hi, j)
		}

	}

}
func partition(a Comparable, lo, hi int) int {
	var pos int
	return pos
}

func sort(data Comparable) {
	n := data.len()
	quickSort(data, 0, n-1)
	//bubbleSort(data)
}

func main() {
	//arr := [...]int{20, 4, 5, 3, 8, 13, 1, 66, 80, 2}
	arr := []int{3, 7, 9, 8, 38, 93, 12, 222, 45, 93, 23, 84, 65, 2}
	//arr := RandArray(10)
	fmt.Println("Initial array is:", arr)
	fmt.Println("")
	//bubbleSort(b[:])
	//selectSort(b[:])
	//insertSort(b[:])
	sort(IntSlice(arr[:]))
	fmt.Println("sorted array is:", arr)

}
