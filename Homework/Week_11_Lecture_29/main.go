package main

import "log"


type Order struct {
	Customer string
	Amount   int
}

func GroupBy[T any, U comparable](col []T, keyFn func(T) U) map[U][]T {
	m := make(map[U][]T)
	for _, v := range col {
		k := keyFn(v)
		m[k] = append(m[k], v)
	}
	return m
}

func main() {
	results := GroupBy([]Order{
		{Customer: "John", Amount: 1000},
		{Customer: "Sara", Amount: 2000},
		{Customer: "Sara", Amount: 1800},
		{Customer: "John", Amount: 1200},
	}, func(o Order) string { return o.Customer })

	log.Println(results)
}
