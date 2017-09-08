package main

import (
	utils "../utils"
	"fmt"
)

type Comparable interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

//先实现这个排序接口
type IntSlice []int

func (is IntSlice) Len() int           { return len(is) }
func (is IntSlice) Less(i, j int) bool { return is[i] < is[j] } //升序
func (is IntSlice) Swap(i, j int)      { is[i], is[j] = is[j], is[i] }

//检查是否有序
func isSorted(data Comparable) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}

//用引用类型【切片】来操作对应数组
//冒泡
//相邻元素比较【互换】
func bubbleSort(a Comparable) {
	var swap int
	for i := 0; i < a.Len(); i++ {
		for j := 1; j < a.Len()-i; j++ {
			if a.Less(j, j-1) {
				a.Swap(j, j-1)
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

func QuickSort(data Comparable) {
	n := data.Len()
	lo, hi := 0, n-1
	quickSort(data, lo, hi)
}

//快排
func quickSort(a Comparable, lo, hi int) {
	if lo < hi {
		i, j, pos, last := lo, hi, lo, 0 //0就是使用第一个作为基准值,
		// last这个变量时为了记住最后一次交换变量时出现在哪次
		for i < j {
			for i < j && a.Less(pos, j) {
				j--
			}
			if i < j {
				a.Swap(i, j)
				i++
				pos = j
				last = 1
			}
			for 1 < j && a.Less(i, pos) { //比pos大的放在后面出现的坑中
				i++
			}
			if i < j {
				a.Swap(i, j)
				j--
				pos = i
				last = -1
			}
		}
		if last == 1 {
			a.Swap(j, pos)
		} else if last == -1 {
			a.Swap(i, pos)
		}
		quickSort(a, lo, i-1)
		quickSort(a, i+1, hi)
	}

}

func main() {
	//arr := [...]int{20, 4, 5, 3, 8, 13, 1, 66, 80, 2}
	arr := utils.RandArray(20)
	//arr = []int{3, 7, 9, 8, 38, 93, 12, 222, 45, 93, 23, 84, 65, 2}

	fmt.Println("Initial array is:", arr)
	fmt.Println("")
	//bubbleSort(b[:])
	//selectSort(b[:])
	//insertSort(b[:])
	t := new(utils.Stopwatch) //封装的计时间的方法
	t.Start()
	QuickSort(IntSlice(arr))
	t.Stop()
	fmt.Println(t.RuntimeMs(), "ms")
	fmt.Println("sorted array is:", arr)

}
