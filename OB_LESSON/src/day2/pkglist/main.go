package pkglist

import (
	"add"
	"fmt"
)

func list(n int) {
	for i := 0; i <= n; i++ {
		fmt.Printf("%d+%d=%d\n", i, n-i, n)
	}
}

func main() {
	list(add.Num1 + add.Num2)
}
