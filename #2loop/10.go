package main

import "fmt"

// My solution
/*
func main() {
	var a, b, c int
	fmt.Scanf("%d %d", &a, &b)
	a, b = max(a, b), min(a, b)
	for a%b != 0 {
		c = a % b
		a = b
		b = c
	}
	fmt.Println("Ekub:", c)
}
*/

// GPT solution
func main() {
	var a, b int
	fmt.Scan(&a, &b)

	for b != 0 {
		a, b = b, a%b
	}

	fmt.Println("GCD:", a)
}
