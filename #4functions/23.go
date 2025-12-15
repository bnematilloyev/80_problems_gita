package main

import "fmt"

func digitCount(x int) int {
	count := 0
	for x > 0 {
		x /= 10
		count++
	}
	return count
}

func main() {
	var x int
	fmt.Scan(&x)
	fmt.Println(digitCount(x))
}
