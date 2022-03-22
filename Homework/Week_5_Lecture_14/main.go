package main

import (
	"context"
	"fmt"
	"time"
)

type BufferedContext struct {
	ctx        context.Context
	timeout    time.Duration
	bufferSize int
}

func NewBufferedContext(timeout time.Duration, bufferSize int) *BufferedContext {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return &BufferedContext{ctx: ctx, timeout: timeout, bufferSize: bufferSize}
	/*Implement the rest */
}

func (bc *BufferedContext) Done() <-chan struct{} {
	return nil
	/* This function will serve in place of the oriignal context */

	/*
		make it so that the result channel gets closed in one of the to cases;
		a) the emebdded context times out
		b) the buffer gets filled
	*/
}

func (bc *BufferedContext) Run(fn func(context.Context, chan string)) {
	/* This function serves for executing the test */
	/* Implement the rest */
}

func main() {
	ctx := NewBufferedContext(time.Second, 10)
	ctx.Run(func(ctx context.Context, buffer chan string) {
		for {
			select {
			case <-ctx.Done():
				return
			case buffer <- "bar":
				time.Sleep(time.Millisecond * 200) // try different values here
				fmt.Println("bar")
			}
		}
	})
}
