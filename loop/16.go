package main

import (
	"fmt"
	"math"
)

func main() {
	var n, x, minInd, maxInd int
	max := math.MinInt
	min := math.MaxInt

	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&x)
		if x >= max {
			max = x
			maxInd = i
		}
		if x <= min {
			min = x
			minInd = i
		}
	}
	if minInd >= maxInd {
		fmt.Println(min, minInd)
	} else {
		fmt.Println(max, maxInd)
	}
}
