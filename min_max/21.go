package main

import "fmt"

func main() {
	var n, x, prev, maxSum int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		if i > 0 {
			maxSum = max(maxSum, prev+x)
		}
		prev = x
	}
	fmt.Println(maxSum)
}
