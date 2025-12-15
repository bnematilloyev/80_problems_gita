package main

import "fmt"

func main() {
	var n, length, sum int
	fmt.Scanf("%d", &n)
	for n > 0 {
		length++
		sum += n % 10
		n = n / 10
	}
	fmt.Println(length, sum)
}
