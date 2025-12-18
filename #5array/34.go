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
	left, right := 0, n-1
	for left < right {
		fmt.Println(arr[left])
		fmt.Println(arr[right])
		left++
		right--
	}
	if n%2 == 1 {
		fmt.Println(arr[left])
	}
}
