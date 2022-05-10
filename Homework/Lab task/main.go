package main

import (
	"log"
	"sync"
	"time"
)

/*
	scheduleTask must be able to execute the given task at regular intervals (specified by interval), until the givne tiemout is reached.
	It is important to make sure that tasks do not overlap one anotehr, i.e. it should not be possible to have more than one running at the same time.
	Also, make sure that the interval is "fixed" - if a task takes less than the interval, then next must start exactly at the start of th enext interval.
	However, if a task takes longer than one interval, you have to make sure that the next one starts only after the previous task has finished.

	Make sure to account for panics in the code, and handel them gracefully, so that the execution continues.
*/

func scheduleTask(task func() error, interval time.Duration, timeout time.Duration) {
	//add waitgroup and mutex
	var wg sync.WaitGroup
	var mtx sync.Mutex
	startTime := time.Now()
	//add each task to a goroutine and set interval and timeout to the function
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			currentTime := time.Now()
			//check if the time is up
			if time.Since(startTime) > timeout {
				break
			}

			//do the task in the given interval
			for time.Since(currentTime) < interval {
				mtx.Lock()
				func() {
					defer func() {
						if r := recover(); r != nil {
							log.Println("Recovering from panic:", r)
						}
					}()

					task()
				}()
				time.Sleep(interval - time.Since(currentTime))
				mtx.Unlock()
			}
		}
	}()
	wg.Wait()
}

func main() {
	var a = 0

	scheduleTask(func() error {
		time.Sleep(50 * time.Millisecond)
		if a == 3 {
			a++
			panic("oops")
		}
		a++
		return nil
	}, time.Millisecond*100, time.Millisecond*1000)

	log.Println(a)
	if a != 10 {
		log.Fatal("Expected it to be 10")
	}
}
