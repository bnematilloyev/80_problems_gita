package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		n   = 0
		min = math.MaxInt32
		max = math.MinInt32
		x   = 0
	)
	fmt.Scan(&n)
	for n > 0 {
		fmt.Scan(&x)
		if x > max {
			max = x
		}
		if x < min {
			min = x
		}
		n -= 1
	}
	fmt.Println(min, max)
}
