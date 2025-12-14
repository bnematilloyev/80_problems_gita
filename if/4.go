package main

import "fmt"

// Ikkita butun son berilgan Day (kun) va Month (oy). (Kabisa bo`lmagan yil sanasi
// kiritiladi). Berilgan sanadan keyingi sanani ifodalovchi programma tuzilsin.
func main() {
	var day, month, month_days int
	fmt.Scanf("%d %d", &day, &month)
	if day > 31 {
		fmt.Println("Bunday sana yo'q")
	}
	switch month {
	case 1:
		month_days = 31
		break
	case 2:
		month_days = 28
		break
	case 3:
		month_days = 31
		break
	case 4:
		month_days = 30
		break
	case 5:
		month_days = 31
		break
	case 6:
		month_days = 30
		break
	case 7:
		month_days = 31
		break
	case 8:
		month_days = 31
		break
	case 9:
		month_days = 30
		break
	case 10:
		month_days = 31
		break
	case 11:
		month_days = 30
		break
	case 12:
		month_days = 31
		break
	default:
		fmt.Println("Bunday oy yo'q")
	}
	if day <= 31 && month <= 12 {
		if month_days == day && month+1 > 12 {
			fmt.Println(1, 1)
		} else if month_days == day {
			fmt.Println(1, month+1)
		} else {
			fmt.Println(day+1, month)
		}
	}
}
