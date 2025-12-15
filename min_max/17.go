package main

import (
	"fmt"
	"math"
)

func main() {
	var n, x int
	minPositive := math.MaxInt
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		if x > 0 {
			minPositive = min(minPositive, x)
		}
	}
	if minPositive < 0 || minPositive == math.MaxInt {
		fmt.Println(0)
	} else {
		fmt.Println(minPositive)
	}
}
