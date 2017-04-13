package main

import (
	"fmt"
)

type USB interface {
	Name() string
	Connecter
}
type Connecter interface {
	Connect()
}
type PhoneConnecter struct {
	name string
}

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
