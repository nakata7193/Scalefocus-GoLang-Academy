package main

import "fmt"

func daysInMonth(month int, year int) (int, bool) {
	switch month {
	case 1:
		return 31, true
		break
	case 2:
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			return 29, true
		} else {
			return 28, true
		}
		break
	case 3:
		return 31, true
		break
	case 4:
		return 30, true
		break
	case 5:
		return 31, true
		break
	case 6:
		return 30, true
		break
	case 7:
		return 31, true
		break
	case 8:
		return 31, true
		break
	case 9:
		return 30, true
		break
	case 10:
		return 31, true
		break
	case 11:
		return 30, true
		break
	case 12:
		return 31, true
		break
	default:
		return 0, false
	}
} //missing return??

func main() {
	fmt.Println(daysInMonth(6, 1992))
}
