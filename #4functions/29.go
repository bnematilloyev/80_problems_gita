package main

import "fmt"

func main() {
	var n, x, prev, GCD int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		if i > 0 {
			GCD = gcd(x, prev)
		}
		prev = x
	}
	fmt.Println(GCD)
}
