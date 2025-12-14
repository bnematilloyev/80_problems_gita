package main

import "fmt"

// Perfect Number
func main() {
	var N int
	fmt.Scanf("%d", &N)
	for i := 4; i <= N; i++ {
		k := 0
		for j := 1; j < i; j++ {
			if i%j == 0 {
				k += j
			}
		}
		if k == i {
			fmt.Println(i)
		}
	}
}
