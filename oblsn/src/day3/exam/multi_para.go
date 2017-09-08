package exam

import "fmt"

func add(a int, b ...int) int {
	var sum int = a
	for i := 0; i < len(b); i++ {
		sum += b[i]
	}
	return sum
}
func main() {
	sum := add(1, 2, 3, 4, 5, 5)
	fmt.Println(sum)
}
