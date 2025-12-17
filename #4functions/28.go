package main

import "fmt"

func GCD3(a int, b int, c int) int {
	return gcd(a, gcd(b, c))
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(GCD3(a, b, c))
}
