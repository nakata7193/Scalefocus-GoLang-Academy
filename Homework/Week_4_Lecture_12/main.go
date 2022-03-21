package main

import (
	"log"
	"time"
)

func generateThrottled(data string, bufferLimit int, clearInterval time.Duration) <-chan string {
	channel := make(chan string, bufferLimit)
	go func() {
		for {
			timeoutChan := time.After(clearInterval)
			select {
			case channel <- data:
				channel <- data
				<-timeoutChan
			case <-timeoutChan:
				<-channel
			}
		}
	}()
	return channel
}

func main() {
	out := generateThrottled("foo", 2, time.Second)
	for f := range out {
		log.Println(f)
	}
}
