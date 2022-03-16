package main

import (
	"context"
	"time"
)

type BufferedContext struct {
	context.Context
	/* Add other fields you might need */
}

func NewBufferedContext(timeout time.Duration, bufferSize int) *BufferedContext {
	ctx, cancel := context.WithTimeout(context.Background(), timeout) /*Implement the rest */
	return nil

}

func (bc *BufferedContext) Done() <-chan struct{} {
	/* This function will serve in place of the oriignal context */
	/* make it so that the result channel gets closed in one of the to cases;
	a) the emebdded context times out        b) the buffer gets filled    */
	return nil
}

func (bc *BufferedContext) Run(fn func(context.Context, chanstring)) {
	/* This function serves for executing the test */ /* Implement the rest */
}

func main() {

}
