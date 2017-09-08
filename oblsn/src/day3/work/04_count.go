package work

import (
	"bufio"
	"fmt"
	"os"
)

func string_count(str string) (wd_cnt, sps_cnt, num_cnt, ot_cnt int) {
	t := []rune(str)
	for _, v := range t {
		switch {
		case v >= 'a' && v <= 'z':
			fallthrough
		case v >= 'A' && v <= 'Z':
			wd_cnt++
		case v == ' ':
			sps_cnt++
		case v >= '0' && v <= '9':
			num_cnt++
		default:
			ot_cnt++

		}
	}
	return
}

func main() {
	read := bufio.NewReader(os.Stdin)
	res, _, err := read.ReadLine()
	if err != nil {
		fmt.Println("read from console:", err)
		return
	}
	//res此时类型是[]uint8 需要强制类型转换为string
	wc, sc, nc, otc := string_count(string(res))
	fmt.Printf("world cnt:%d\n space cnt:%d\n number cnt:%d\n other cnt:%d\n", wc, sc, nc, otc)

}
