package main

import "fmt"

func DigitN(k int, n int) int {
	i := 1
	for k > 0 {
		d := k % 10
		k /= 10
		i++
		if i == n {
			return d
		}
	}
	return -1
}

func main() {
	var k, n int
	fmt.Scan(&k, &n)
	fmt.Println(DigitN(k, n))
}
