package main

import (
	"fmt"
)

func map_test1() {
	//定义的同时 初始化
	var a map[string]string = map[string]string{
		"key1": "value1",
	}
	a["setting1"] = "yes"
	a["setting2"] = "disable"
	fmt.Println(a)
}
func map_test2() {
	//定义的时候不初始化 只定义第二层map长度
	a := make(map[string]map[string]string, 100)
	a["key1"] = make(map[string]string)
	a["key1"]["key2"] = "abc"
	a["key1"]["key3"] = "abc"
	a["key1"]["key4"] = "abc"
	a["key1"]["key5"] = "abc"
	fmt.Println(a)
}
func modify(a map[string]map[string]string) {
	//查询键值,若不存在则初始化一只
	_, ok := a["Shanshan"]
	if !ok {
		a["Shanshan"] = make(map[string]string)
	}

	a["Shanshan"]["passwd"] = "5201314"
	a["Shanshan"]["nickname"] = "33"
	return
}
func map_test3() {
	a := make(map[string]map[string]string, 50)
	fmt.Println(a)
	modify(a)
	fmt.Println(a)
}

//解析
func trans(m map[string]map[string]string) {
	for key1, var1 := range m {
		fmt.Println(key1)
		for key2, var2 := range var1 {
			fmt.Println("\t", key2, var2)
		}
	}
}

func map_test4() {
	//嵌套map的定义
	a := make(map[string]map[string]string, 100)
	//嵌套map的初始化
	a["key1"] = make(map[string]string)
	a["key1"]["key2"] = "abc"
	a["key1"]["key3"] = "abc"
	a["key1"]["key4"] = "abc"
	a["key1"]["key5"] = "abc"

	a["key2"] = make(map[string]string)
	a["key2"]["key2"] = "abc"
	a["key2"]["key3"] = "abc"

	//解析一波
	trans(a)
	//删除key1
	delete(a, "key1")
	fmt.Println("after tarns:")
	trans(a)

	fmt.Println(len(a))

}

//slice of map
func map_test5() {
	var a []map[int]int
	a = make([]map[int]int, 6)
	/*WAY 1*/
	// if a[0] == nil {
	// 	a[0] = make(map[int]int)
	// }
	// if a[1] == nil {
	// 	a[1] = make(map[int]int)
	// }
	/* WAY 2*/
	// for _, v := range a {
	// 	if v == nil {
	// 		v = make(map[int]int)
	// 	}
	// }
	/*WAY 3*/
	for i := 0; i < len(a); i++ {
		if a[i] == nil {
			a[i] = make(map[int]int)
		}
	}
	a[0][4] = 10
	a[1][5] = 2
	a[1][3] = 3
	fmt.Println(a)
}

func main() {
	map_test5()
}
