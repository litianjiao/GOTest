package main

import (
	//"fmt"
	serial"github.com/tarm/goserial"
	"log"
)
//二分法查找中值
func binaryseach(iArray[10] int64, key int64, n int64)  (iMid   int64) {

	var iLow int64 = 0
	var iHigh= n - 1
	for (iLow <= iHigh) {
		iMid = (iHigh + iLow) / 2
		if (iArray[iMid] > key) {
			iHigh = iMid - 1
		} else if iArray[iMid] < key {
			iLow = iMid + 1
		} else {
			return iMid
		}

	}
return -1
}

//最简单的串口交互
func main() {

		c := &serial.Config{Name: "COM10", Baud: 115200}
		s, err := serial.OpenPort(c)
		if err != nil {
			log.Fatal(err)
		}

		n, err := s.Write([]byte("test"))
		if err != nil {
			log.Fatal(err)
		}

		buf := make([]byte, 128)
		n, err = s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Print("%q", buf[:n])
	}
