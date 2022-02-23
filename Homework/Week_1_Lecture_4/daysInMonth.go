package main

import "fmt"

func daysInMonth(month int, year int) (int, bool) {
	switch month {
	case 1:
		return 31, true
	case 2:
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			return 29, true
		} else {
			return 28, true
		}
	case 3:
		return 31, true
	case 4:
		return 30, true
	case 5:
		return 31, true
	case 6:
		return 30, true
	case 7:
		return 31, true
	case 8:
		return 31, true
	case 9:
		return 30, true
	case 10:
		return 31, true
	case 11:
		return 30, true
	case 12:
		return 31, true
	default:
		return 0, false 
	}
}

func main() {
	fmt.Print(daysInMonth(16, 2020))
}
