package main

import (
	"fmt"
	"math"
)

func main() {
	var n, min_ind int
	fmt.Scan(&n)
	min := math.MaxInt
	for i := 1; i <= n; i++ {
		var x int
		fmt.Scan(&x)
		if x < min {
			min = x
			min_ind = i
		}
	}
	fmt.Println(min, min_ind)
}
