package main

import (
	"fmt"
)

func processOdd(inputs []int) chan int {
	oddValues := make(chan int) //output channel for the value
	for numbers := range inputs {
		go func(num int) {
			if num%2 != 0 {
				oddValues <- num
			}
		}(numbers)
	}
	return oddValues
}

func processEven(inputs []int) chan int {
	evenValues := make(chan int) //output channel for the values
	for numbers := range inputs {
		go func(num int) {
			if num%2 == 0 {
				evenValues <- num
			}
		}(numbers)
	}
	return evenValues
}

func main() {
	inputs := []int{1, 17, 34, 56, 2, 8}
	evenCH := processEven(inputs)
	for val := range evenCH {
		fmt.Print(val)
	}
}
