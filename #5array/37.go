package main

import (
	"fmt"
	"math"
)

func main() {
	var n, R, x int
	fmt.Scan(&n, &R)
	closest := 0
	minDiff := math.MaxInt
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		diff := int(math.Abs(float64(x - R)))
		if minDiff > diff {
			minDiff = diff
			closest = x
		}
	}
	fmt.Println(closest)
}
