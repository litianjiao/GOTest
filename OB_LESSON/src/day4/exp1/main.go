package main

import "fmt"

// new：用来分配内存，主要用来分配值类型，比如int、struct。返回的是指针
// make：用来分配内存，主要用来分配引用类型，比如chan、map、slice
func test(){

	//new return type point have NOT init it
	s1:=new([]int)
	fmt.Println(s1)
	//make return val
	s2:=make([]int, 10)
	fmt.Println(s2)
	//use make init the new obj 
	*s1=make([]int, 5)
	(*s1)[0]=100
	s2[0]=200
	*s1=append(*s1,100,100)
	s2=append(s2,100,100)
	fmt.Println(s1)
	fmt.Println(s2)
    return
}


func main(){
	test()
}