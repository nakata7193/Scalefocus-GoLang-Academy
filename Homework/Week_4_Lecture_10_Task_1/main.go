package main

import (
	"fmt"
	"sync"
)

type ConcurrentPrinter struct {
	sync.WaitGroup
	sync.Mutex
}

func (cp *ConcurrentPrinter) printFoo(times int) {
	cp.WaitGroup.Add(1)
	defer cp.WaitGroup.Done()
	cp.Lock()
	go func() {
		fmt.Print("|foo|")
		cp.Unlock()
	}()
}
func (cp *ConcurrentPrinter) printBar(times int) {
	cp.WaitGroup.Add(1)
	defer cp.WaitGroup.Done()
	cp.Lock()
	go func() {
		fmt.Print("|bar|")
		cp.Unlock()
	}()
}

func main() {
	times := 10
	cp := &ConcurrentPrinter{}
	for i := 0; i <= times; i++ {
		cp.printFoo(i)
		cp.printBar(i)
	}
	cp.WaitGroup.Wait()
}
