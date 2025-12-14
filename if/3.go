package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	if a < b {
		a, b = (a+b)/2, a*b*2
	} else if a > b {
		a, b = a*b*2, (a+b)/2
	}
	fmt.Println(a, b)
}
