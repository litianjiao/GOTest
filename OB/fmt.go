package main

import "fmt"

var g_u8 byte
var g_b bool
var g_int int
var g_f float32
var g_str string

func main() {
	g_u8 = 0x11
	g_b = true
	g_int = 20
	g_f = 0.0000455
	g_str = "hello"
	fmt.Println(g_str)
	fmt.Println(g_b)
	fmt.Println(g_int)
	fmt.Printf("%x\n", g_u8)
	fmt.Println(g_f)

}
