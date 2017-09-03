package main

import (
	"fmt"
)

type Car struct {
	weight int
	name   string
}

type Bike struct {
	Car
	wheel int
}
type Train struct {
	c Car
}

func (p *Car) Run() {
	fmt.Println(p.name, "is running")
}

func main() {
	var a Bike
	a.weight = 100
	a.wheel = 2
	a.name = "捷安特"
	fmt.Println(a)
	a.Run()
	var b Train
	b.c.name = "CRH"
	b.c.weight = 5000
	b.c.Run()
}
