package main

import (
	"log"
	"time"
)

func generateThrottled(data string, bufferLimit int, clearInterval time.Duration) <-chan string {
	channel := make(chan string, bufferLimit)
	handleChannel := make(chan string)
	go func() {
		for {
			handleChannel <- data
			select {
			case n := <-handleChannel:
				channel <- n
			case <-timeAfter(clearInterval,channel):

		}
	}()
	return channel
}

func timeAfter(after time.Duration,ch chan string) chan time.Time {
	done := make(chan time.Time)
	go func() {
		time.Sleep(after)
		
	}()
	return done
}

func main() {
	out := generateThrottled("foo", 2, time.Second)
	for f := range out {
		log.Println(f)
	}
}
