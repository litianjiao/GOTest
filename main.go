package main

import (
	"fmt"
)

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


func main() {
	a:=[10] int64{1,2,3,4,5,6,7,8,9,10}
   fmt.Println("the position is ",(binaryseach(a,4,10)+1))


}