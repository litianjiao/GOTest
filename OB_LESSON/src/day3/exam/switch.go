package exam

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n int
	rand.Seed(time.Now().UnixNano())
	n = rand.Intn(100)
	fmt.Println("rand number is created success,u can guess it")
	for {
		var input int
		fmt.Scanf("%d\n", &input)
		f_quit := false
		switch {
		case input == n:
			fmt.Println("u r right!")
			f_quit = true
		case input > n:
			fmt.Println("large")
		case input < n:
			fmt.Println("small")
		}
		if f_quit {
			break
		}
	}
}
