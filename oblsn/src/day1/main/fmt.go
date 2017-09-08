package main

import "fmt"

var g_u8 byte
var g_b bool
var g_int int
var g_f float32
var g_str string

/*
整数：
老师原意是同一个整数用多种格式打印
%b	二进制表示
%c	相应Unicode码点所表示的字符
%d	十进制表示
%o	八进制表示
%q	单引号围绕的字符字面值，由Go语法安全地转义
%x	十六进制表示，字母形式为小写 a-f
%X	十六进制表示，字母形式为大写 A-F
%U	Unicode格式：U+1234，等同于 "U+%04X"

*/
func main() {
	g_u8 = 0x11
	g_b = true
	g_int = 20
	g_f = 0.0000455
	g_str = "hello"
	fmt.Println(g_str)
	fmt.Println(g_b)
	fmt.Println(g_int)
	fmt.Printf("%x\n", g_u8) //hex
	fmt.Printf("%b\n", g_u8) //bin
	fmt.Printf("%v\n", g_u8) //default
	fmt.Println(g_f)

	//int2str
	str := fmt.Sprintf("%d", g_u8)
	fmt.Printf("%s\n", str)
	fmt.Printf("%q", str) //原字符串内容打印

}
