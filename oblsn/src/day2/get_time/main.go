package get_time

import (
	"fmt"
	"time"
)

const (
	Man    = 1
	Female = 2
)

func main() {
	for {
		second := time.Now().Second()
		if second%Female == 0 {
			fmt.Println("Female")
		} else {
			fmt.Println("Man")
		}
		time.Sleep(time.Second)

	}
}
