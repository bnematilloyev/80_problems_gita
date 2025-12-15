package main

import (
	"fmt"
	"math"
)

func main() {
	var n, x, min_ind, max_ind int
	max := math.MinInt
	min := math.MaxInt

	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&x)
		if x >= max {
			max = x
			max_ind = i
		}
		if x < min {
			min = x
			min_ind = i
		}
	}
	fmt.Println(min, min_ind)
	fmt.Println(max, max_ind)
}
