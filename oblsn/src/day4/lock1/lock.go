package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var lock sync.Mutex     //互斥锁 性能较低 但稳定
var rwlock sync.RWMutex //读写锁 用于读大于写

func testLock() {
	var a map[int]int
	a = make(map[int]int)
	a[8] = 10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[18] = 10
	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			lock.Lock()
			b[8] = rand.Intn(100)
		}(a)
	}
	lock.Lock()
	fmt.Println(a)
	lock.Unlock()
	time.Sleep(time.Second)

}

func testRWLock() {
	var a map[int]int
	a = make(map[int]int, 6)
	var count int32
	a[8] = 10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[18] = 10

	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			lock.Lock()
			b[8] = rand.Intn(100)
			time.Sleep(10 * time.Millisecond)
			lock.Unlock()
		}(a)
	}
	//在这里建立了100个协程来不断读取map a
	//通过打印最终原子数发现读写锁处理次数大大多于互斥锁
	for i := 0; i < 100; i++ {
		go func(b map[int]int) {
			for {
				//lock.Lock()
				rwlock.RLock()
				time.Sleep(time.Millisecond)
				//fmt.Println(a)
				rwlock.RUnlock()
				//lock.Unlock()
				atomic.AddInt32(&count, 1) //原子数自加
			}
		}(a)
	}
	time.Sleep(time.Second * 3)
	fmt.Println(atomic.LoadInt32(&count)) //原子数读取
}

func main() {
	//testLock()
	testRWLock()
}
