package main

import (
	"fmt"
	"sort"
	"time"
)

type TimeSlice []time.Time

func sortDates(format string, dates ...string) ([]time.Time, error) {
	var sortedDates []time.Time
	for _, date := range dates {
		date, err := time.Parse(format, date)

		if err != nil {
			fmt.Println(err)
		} else {
			sortedDates = append(sortedDates, date)
		}
	}

	sort.Slice(sortedDates, func(i, j int) bool {
		return sortedDates[i].Before(sortedDates[j])
	})

	return sortedDates, nil
}

func main() {
	dates := []string{"Sep-14-2008", "Dec-03-2021", "Mar-18-2022"}
	format := "Jan-02-2006"
	fmt.Print(sortDates(format, dates...))
}
