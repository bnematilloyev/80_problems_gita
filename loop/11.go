package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	butun := 0
	qoldiq := a

	for qoldiq >= b {
		qoldiq -= b
		butun++
	}
	fmt.Println(qoldiq, butun)
}
