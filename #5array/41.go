package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	for i := 0; i < n-1; i++ {
		minInd := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minInd] {
				minInd = j
			}
		}
		if minInd != i {
			arr[i], arr[minInd] = arr[minInd], arr[i]
		}
	}
	fmt.Print(arr, "\n")
}
