package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n)
	arr := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		arr[x]++
	}

	counts, items := make([]int, len(arr)), make([]int, len(arr))
	i := 0
	for count, item := range arr {
		counts[i] = count
		items[i] = item
		i++
	}
	fmt.Print(counts, "\n")
	fmt.Print(items, "\n")
}
