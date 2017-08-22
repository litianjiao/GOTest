package main

import "fmt"
//Modify the specified number with closure
func modifyNum(editnum int)func( int) int {
	return func(source int)int{
		if(source==0){
			source=editnum
		}
		return source
	}
}

func main()  {
	var a [10]int
	//a[0]=10
	for i,v:=range a{
		fmt.Printf("a[%d]=%d\n",i,v)
	}
	modi10:=modifyNum(10)
	modi20:=modifyNum(20)
	for i,v:=range a{
		fmt.Printf("modify10 a[%d]=%d \n",i,modi10(v))
	}
	for i,v:=range a{
		fmt.Printf("modify20 a[%d]=%d \n",i,modi20(v))
	}
	fmt.Println(a)
}