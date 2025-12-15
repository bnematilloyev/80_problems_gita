package main

import "fmt"

func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n, x, primesCount int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		if isPrime(x) {
			primesCount++
		}
	}
	fmt.Println(primesCount)
}
