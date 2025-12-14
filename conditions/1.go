package main

import (
	"fmt"
)

// Uchta son berilgan. Shu sonlarning kichigini
// aniqlovchi programma tuzilsin.

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	minum := a
	if b < minum {
		minum = b
	}
	if c < minum {
		minum = c
	}
	fmt.Println(minum)
	fmt.Println(min(a, b, c))
}
