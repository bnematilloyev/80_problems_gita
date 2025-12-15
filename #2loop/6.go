package main

import "fmt"
import "math"

func main() {
	var N int
	fmt.Scanf("%d", &N)
	for i := 2; i <= N; i++ {
		root := int(math.Sqrt(float64(i)))
		isPrime := true
		for j := 2; j <= root; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(i)
		}
	}
}
