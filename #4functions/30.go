package main

import "fmt"

func isLeapYear(Y int) bool {
	if Y%4 == 0 {
		if Y%100 == 0 {
			return Y%400 == 0
		}
		return true
	}
	return false
}

func main() {
	var x, count int
	for i := 0; i < 3; i++ {
		fmt.Scan(&x)
		if isLeapYear(x) == true {
			count++
		}
	}
	fmt.Println(count)
}
