package utils

import (
	"math/rand"
	"time"
)

type Stopwatch struct {
	start time.Time
	stop  time.Time
}

func (s *Stopwatch) Start() {
	s.start = time.Now()
}

func (s *Stopwatch) Stop() {
	s.stop = time.Now()
}

//纳秒

func (s Stopwatch) RuntimeNs() int {
	return s.stop.Nanosecond() - s.start.Nanosecond()
}

//微秒
func (s Stopwatch) RuntimeUs() float64 {
	return (float64)(s.stop.Nanosecond()-s.start.Nanosecond()) / 1000.00
}

//毫秒
func (s Stopwatch) RuntimeMs() float64 {
	return (float64)(s.stop.Nanosecond()-s.start.Nanosecond()) / 1000000.00
}

//秒
func (s Stopwatch) RuntimeS() float64 {
	return (float64)(s.stop.Nanosecond()-s.start.Nanosecond()) / 10000000000.00
}

func RandArray(n int) []int {
	arr := make([]int, n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i <= n-1; i++ {
		arr = append(arr, r.Intn(100))
	}
	return arr
}
