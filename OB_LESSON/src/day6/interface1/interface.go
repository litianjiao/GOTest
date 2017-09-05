package main

import "fmt"

type Passenger struct {
	rideCar string
}
type EVCar interface {
	Car
	charge(int)
}

type LuxuryCar interface {
	Car
	PickUpGirl()
}

type Car interface {
	GetName() string
	GetPower() string
	Run()
	Horn()
	//	AddPassenger(int)
}
type BENZ struct {
	Name  string
	Power string
	Hezai int
}
type BYD struct {
	Name    string
	Power   string
	Battery string
}

func (b *BENZ) GetName() string {
	return b.Name
}
func (b *BENZ) GetPower() string {
	return b.Power
}
func (b *BENZ) Run() {
	fmt.Println(b.Name, "is running")
}
func (b *BENZ) Horn() {
	fmt.Println("DIDIDI")
}
func (b *BENZ) PickUpGirl() {
	fmt.Println("piapiapia")
}

func (b *BYD) GetName() string {
	return b.Name
}
func (b *BYD) GetPower() string {
	return b.Power
}
func (b *BYD) Run() {
	fmt.Println(b.Name, "is running")
}
func (b *BYD) Horn() {
	fmt.Println("DIDIDI")
}
func (b *BYD) charge(power int) {

}

func main() {
	var car Car
	var NiceCar LuxuryCar
	fmt.Println(car)
	benz := &BENZ{
		Name:  "BENZ-S400",
		Power: "3.0L",
		Hezai: 5,
	}

	car = benz
	car.Run()
	NiceCar = benz
	NiceCar.Horn()
	NiceCar.PickUpGirl()

	byd := &BYD{
		Name:    "秦",
		Power:   "1.5T",
		Battery: "26Ah",
	}
	car = byd
	car.Run()
}
