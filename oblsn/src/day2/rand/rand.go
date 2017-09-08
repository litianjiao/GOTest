package rand

import (
	"fmt"
	"math/rand"
	"time"
)

/*
******************************************************************
  * @brief     rand
  * @param
  * @ret
  * @author    TRoy
  * @date      2017/7/3 21:43
******************************************************************
*/
//init函数属于保留函数 同一个包内可定义多个 执行顺序自上而下在main函数执行之前调用

func init() {
	rand.Seed(time.Now().UnixNano())
}
func init() {
	rand.Seed(1)
}

func main() {
	rand.Seed(2)
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Println(a)

	}
	for i := 0; i < 10; i++ {
		a := rand.Intn(100)
		fmt.Println(a)
	}
	for i := 0; i < 10; i++ {
		a := rand.Float32()
		fmt.Println(a)
	}
}
