package main

import (
	"fmt"
)

//USB接口
type USB interface {
	Name() string
	Connecter
}

//连接口
type Connecter interface {
	Connect()
}

//手机连接对象
type PhoneConnecter struct {
	name string
}

//-实现手机上USB口方法（根据interface规定他的名字，实现connecter接口下的connect方法）
func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("Connect:", pc.name)
}
func Disconnect(usb USB) {
	if pc, ok := usb.(PhoneConnecter); ok {
		fmt.Println("Disconnect:", pc.name)
		return
	}
}
func main() {
	//定义对象 手机USB口
	a := PhoneConnecter{"PhoneConnecter"}
	a.Connect()
	Disconnect(a)

}

//type NUM int
//func (i *NUM) Increase(b int) {
//  *i+=NUM(b)
//}
//var a NUM
//a.Increase(100)
//fmt.Println(a)
