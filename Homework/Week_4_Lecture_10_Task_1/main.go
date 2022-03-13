package main

import "sync"

type ConcurrentPrinter struct {
	sync.WaitGroup
	sync.Mutex
}

func (cp *ConcurrentPrinter) printFoo(times int) {

}
func (cp *ConcurrentPrinter) printBar(times int) {

}

func main() {
	times := 10
	cp = &ConcurrentPrinter{}
	cp.PrintFoo(times)
	cp.PrintBar(times)
	cp.Wait()
}
