package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	left, right := 0, n-1
	for left+1 < right-1 {
		fmt.Println(arr[left], arr[left+1])
		fmt.Println(arr[right], arr[right-1])
		left += 2
		right -= 2
	}
	if left == right {
		fmt.Println(arr[left])
	} else if left < right {
		fmt.Println(arr[left], arr[right])
	}
}
