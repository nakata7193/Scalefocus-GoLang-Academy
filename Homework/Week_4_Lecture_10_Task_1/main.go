package main

import (
	"fmt"
	"sync"
	"time"
)

type ConcurrentPrinter struct {
	sync.WaitGroup
	sync.Mutex
	counter     int
	lastElement string
}

func (cp *ConcurrentPrinter) printFoo(times int) {
	cp.Add(1)
	go func() {
		defer cp.Done()
		for {
			if cp.counter == times {
				break
			}
			cp.Lock()
			if cp.lastElement != "foo" {
				fmt.Print("foo")
				cp.lastElement = "foo"
				cp.counter++
			}
			cp.Unlock()
			time.Sleep(10 * time.Millisecond)
		}
	}()
}
func (cp *ConcurrentPrinter) printBar(times int) {
	cp.Add(1)
	go func() {
		defer cp.Done()
		for {
			if cp.counter == times {
				break
			}
			cp.Lock()
			if cp.lastElement != "bar" {
				fmt.Print("bar")
				cp.lastElement = "bar"
				cp.counter++
			}
			cp.Unlock()
			time.Sleep(10 * time.Millisecond)
		}
	}()
}

func main() {
	times := 10
	cp := &ConcurrentPrinter{}
	cp.printFoo(times)
	cp.printBar(times)
	cp.Wait()
}
