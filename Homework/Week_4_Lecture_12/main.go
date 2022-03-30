package main

import (
	"log"
	"time"
)

func generateThrottled(data string, bufferLimit int, clearInterval time.Duration) <-chan string {
	channel := make(chan string, bufferLimit-1)
	go func() {
		timeoutChan := time.After(clearInterval)
		for {
			select {
			case channel <- data:
			default:
				<-timeoutChan
				timeoutChan = time.After(clearInterval)
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
