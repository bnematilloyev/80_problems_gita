package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}

	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(n, func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Print("\n")
}
