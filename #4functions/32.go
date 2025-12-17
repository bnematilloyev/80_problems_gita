package main

import "fmt"

func prevDate(d int, m int, y int) {
	if d > daysCountInMonth(m, y) {
		fmt.Println("Bunday sana yo'q")
	} else if d == 1 && m == 1 {
		fmt.Println(31, 12, y-1)
	} else if d == 1 {
		fmt.Println(daysCountInMonth(m-1, y), m-1, y)
	} else {
		fmt.Println(d-1, m, y)
	}
}

func main() {
	var d, m, y int
	fmt.Scan(&d, &m, &y)
	prevDate(d, m, y)
}
