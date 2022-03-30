package main

import (
	"log"
	"sync"
	"time"
)

func primesAndSleep(n int, sleep time.Duration) []int {
	res := []int{}
	for k := 2; k < n; k++ {
		for i := 2; i < n; i++ {
			if k%i == 0 {
				time.Sleep(sleep)
				if k == i {
					res = append(res, k)
				}
				break
			}
		}
	}
	return res
}

func goPrimesAndSleep(n int, sleep time.Duration) []int {
	var wg sync.WaitGroup
	result := []int{}
	for k := 2; k < n; k++ {
		wg.Add(1)
		go func(k int) {
			for i := 2; i < n; i++ {
				if k%i == 0 {
					time.Sleep(sleep)
					if k == i {
						result = append(result, k)
					}
					wg.Done()
					break
				}
			}
		}(k)
		wg.Wait()
	}
	return result
}

func main() {
	log.Println(primesAndSleep(50, 3*time.Millisecond))
	log.Println(goPrimesAndSleep(50, 3*time.Millisecond))
}
