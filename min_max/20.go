package main

import (
	"fmt"
	"math"
)

func main() {
	var n, x int
	firstMin := math.MaxInt
	secondMin := math.MaxInt
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		if x < firstMin {
			secondMin = firstMin
			firstMin = x
		}
		if x < secondMin && x != firstMin {
			secondMin = x
		}
	}
	fmt.Println(firstMin, secondMin)
}
