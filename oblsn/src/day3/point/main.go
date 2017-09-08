package point

import "fmt"

func modify(p *int) {
	fmt.Println(p)
	*p = 10000000
}

func main() {
	var a int = 10
	fmt.Println(&a)
	var p *int
	p = &a
	fmt.Println(*p)
	*p = 20
	fmt.Println(a)

	var b = 30
	p = &b
	fmt.Println(a)
	fmt.Println(b)

	modify(&a)
	fmt.Println(a)
}
