package main

import "fmt"

func isPalindrom(N int) bool {
	length := digitCount(N)
	i, half := 0, 0
	for i != length/2 {
		d := N % 10
		N /= 10
		half = half*10 + d
		i++
	}
	if length%2 == 1 {
		return half == N/10
	}
	return half == N
}

func main() {
	var n, x, cnt int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		if isPalindrom(x) {
			cnt++
		}
	}
	fmt.Println(cnt)
}
