package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	minMaxLocal := math.MaxInt
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	for i := 1; i < n-1; i++ {
		if arr[i-1] < arr[i] && arr[i] > arr[i+1] {
			minMaxLocal = min(minMaxLocal, arr[i])
		}
	}
	fmt.Println(minMaxLocal)
}
