package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n)
	frequency := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		frequency[x]++
	}
	for elem, count := range frequency {
		fmt.Println(elem, ":", count)
	}
}
