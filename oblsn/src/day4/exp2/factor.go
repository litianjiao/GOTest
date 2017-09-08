package main

import "fmt"
// factorial 
/*
func factor(i int)int{
	if(i==1){
		return 1
	}
	return factor(i-1)*i
}

func main(){
    fmt.Println( factor(5))
}
*/
// feb 
func feb(n int)int{
  if n<=1 {
	  return 1
  }
  return feb(n-1)+feb(n-2)
}

func main(){
//    for i:=0 ;i<10 ;i++{
// 	   n:=feb(i)
//        fmt.Println(n)
//    }
  
  fmt.Println(feb(0))
}