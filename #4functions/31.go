package main

func daysCountInMonth(month int, year int) int {
	switch month {
	case 1:
		return 31
	case 2:
		if isLeapYear(year) == true {
			return 29
		}
		return 28
	case 3:
		return 31
	case 4:
		return 30
	case 5:
		return 31
	case 6:
		return 30
	case 7:
		return 31
	case 8:
		return 31
	case 9:
		return 30
	case 10:
		return 31
	case 11:
		return 30
	case 12:
		return 31
	}
	return 0
}

//func main() {
//	var year, m int
//	fmt.Scan(&year)
//	for i := 0; i < 3; i++ {
//		fmt.Scan(&m)
//		fmt.Println(daysCountInMonth(m, year))
//	}
//}
