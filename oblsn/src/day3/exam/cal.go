package exam

import "fmt"

type op_func func(int, int) int

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func multiply(a, b int) int {
	return a * b
}
func devide(a, b int) int {
	return a / b
}
func operator(op op_func, a int, b int) int {
	return op(a, b)
}
func main() {
	c := add
	fmt.Println(c)
	sum := operator(c, 1, 3)
	fmt.Println(sum)

}
