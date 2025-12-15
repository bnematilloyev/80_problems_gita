package main

import "fmt"

func main() {
	var n, x, maxInd, maxNum int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&x)
		if x >= maxNum {
			maxInd = i
			maxNum = x
		}
	}
	fmt.Println(maxNum, n-maxInd)
}
