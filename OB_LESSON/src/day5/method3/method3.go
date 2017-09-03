package main

import (
	"fmt"
)

type Car struct {
	weight int
	name   string
}

func (p *Car) Run() {
	fmt.Println("running")
}

type Bike struct {
	Car
	wheel int
}

type Train struct {
	Car
}

//特定格式化输出
func (p *Train) String() string {
	str := fmt.Sprintf("name =[%s] weight =[%d]", p.name, p.weight)
	return str
}

func main() {
	var a Bike
	a.weight = 100
	a.name = "Bike1"
	a.wheel = 2
	fmt.Println(a)
	a.Run()

	var b Train
	b.weight = 500
	b.name = "CRH"
	b.Run()
	fmt.Printf("%s\n", b.String())
	//等价
	fmt.Printf("%s\n", &b)
}
