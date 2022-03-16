package main

import (
	"fmt"
	"sync"
)

// func processOdd(inputs []int) chan int {
// 	oddValues := make(chan int, len(inputs))
// 	var wg sync.WaitGroup
// 	for _, numbers := range inputs {
// 		go func(num int) {
// 			wg.Add(1)
// 			if num%2 == 0 {
// 				oddValues <- num
// 			}
// 			wg.Done()
// 		}(numbers)
// 	}
// 	wg.Wait()
// 	close(oddValues)
// 	return oddValues
// }

func processEven(inputs []int) chan int {
	evenValues := make(chan int, len(inputs))
	var wg sync.WaitGroup
	for _, numbers := range inputs {
		go func(num int) {
			wg.Add(1)
			if num%2 == 0 {
				evenValues <- num
			}
			wg.Done()
		}(numbers)
	}
	wg.Wait()
	close(evenValues)
	return evenValues
}

func main() {
	inputs := []int{1, 17, 34, 56, 2, 8}
	evenCH := processEven(inputs)
	for val := range evenCH {
		fmt.Print(val)
	}
}
