package main

import (
	"fmt"
)

func main() {
	var x, p, final float64
	var month int
	fmt.Scanf("%f %f", &x, &p)
	final = x * 2
	for x < final {
		month++
		x += x * p / 100
	}
	val := fmt.Sprintf("%.2f", x)
	fmt.Println(month, "oyda", val, "so'm")
}
