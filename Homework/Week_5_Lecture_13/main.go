package main

import (
	"context"
	"fmt"
	"time"
)

type BufferedContext struct {
	context.Context
	buffer chan string
	context.CancelFunc
}

func NewBufferedContext(timeout time.Duration, bufferSize int) *BufferedContext {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	buffer := make(chan string, bufferSize)
	newBufferCTX := &BufferedContext{Context: ctx, buffer: buffer, CancelFunc: cancel}
	return newBufferCTX
}

func (bc *BufferedContext) Done() <-chan struct{} {
	if len(bc.buffer) == cap(bc.buffer) {
		bc.CancelFunc()
	}
	return bc.Context.Done()
}

func (bc *BufferedContext) Run(fn func(context.Context, chan string)) {
	fn(bc, bc.buffer)
}

func main() {
	ctx := NewBufferedContext(time.Second, 10)
	ctx.Run(func(ctx context.Context, buffer chan string) {
		for {
			select {
			case <-ctx.Done():
				return
			case buffer <- "bar":
				time.Sleep(time.Millisecond * 1) // try different values here
				fmt.Println("bar")
			}
		}
	})
}
