package main

import "fmt"

func nextDate(d int, m int, y int) {
	if d > daysCountInMonth(m, y) {
		fmt.Println("Bunday sana yo'q")
	} else if d == daysCountInMonth(m, y) && m == 12 {
		fmt.Println(1, 1, y+1)
	} else if d == daysCountInMonth(m, y) {
		fmt.Println(1, m+1, y)
	} else {
		fmt.Println(d+1, m, y)
	}
}

func main() {
	var d, m, y int
	fmt.Scan(&d, &m, &y)
	nextDate(d, m, y)
}
