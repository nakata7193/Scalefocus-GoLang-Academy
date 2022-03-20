package main

import (
	"fmt"
	"sync"
)

func processOdd(inputs []int) chan int {
	oddValues := make(chan int)
	var wg sync.WaitGroup
	for _, num := range inputs {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			if num%2 == 0 {
				oddValues <- num
			}
		}(num)
	}
	go func() {
		wg.Wait()
		close(oddValues)
	}()
	return oddValues
}

func processEven(inputs []int) chan int {
	evenValues := make(chan int)
	var wg sync.WaitGroup
	for _, num := range inputs {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			if num%2 == 0 {
				evenValues <- num
			}
		}(num)
	}
	go func() {
		wg.Wait()
		close(evenValues)
	}()

	return evenValues
}

func main() {
	inputs := []int{1, 17, 34, 56, 2, 8}
	evenCH := processEven(inputs)
	for range inputs {
		fmt.Println(<-evenCH)
	}
}
