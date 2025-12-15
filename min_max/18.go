package main

import "fmt"

func main() {
	var n, x, maxOdd, oddInd int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&x)
		if x%2 == 1 && x > maxOdd {
			maxOdd = x
			oddInd = i
		}
	}
	if maxOdd == 0 {
		fmt.Println(0)
	} else {
		fmt.Println(maxOdd, oddInd)
	}
}
