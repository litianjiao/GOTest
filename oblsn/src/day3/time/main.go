package time

import (
	"fmt"
	//"strings"
	"time"
)

func test() {
	//	s := []string{"foo", "bar", "baz"}
	//	fmt.Println(strings.Join(s, ", "))
	time.Sleep(time.Millisecond * 100)
	//s1 = strings.Join(s1, "66")
}
func main() {
	now := time.Now()
	fmt.Println(now.Format("2006年01月02日 15:04:05"))
	start := time.Now().UnixNano()
	//	s1 := "666555444"
	test()
	end := time.Now().UnixNano()
	fmt.Printf("cost:%d\n", (end-start)/1000)
}
