package exp4_open

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./test.log")
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("read string failed,err:", err)
		return
	}
	fmt.Printf("read str succ,ret:\n", str)
}
