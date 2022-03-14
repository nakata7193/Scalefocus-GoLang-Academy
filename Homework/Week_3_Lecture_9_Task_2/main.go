package main

import (
	"errors"
	"fmt"
	"log"
)

type Action func() error

func SafeExecError(errorFunc Action) Action {
	errorFunc = func() error {
		return errors.New("wrapped error")
	}

	defer func() error {
		if recover := recover(); recover != nil {
			return fmt.Errorf("safe exec: %w", errorFunc())
		}
		return fmt.Errorf("couldn't recover")
	}()
	return errorFunc
}

func SafeExecPanic(errorFunc Action) Action {
	defer func() {
		if recover := recover(); recover != nil {
			fmt.Println("recovered successfully")
		} else {
			panic("crash")
		}
	}()
	return errorFunc
}

func SafeExec(a Action) Action {
	defer func() error {
		return nil
	}()
	return nil
}

func main() {
	var errorFunc Action

	// err := SafeExecError(errorFunc)
	// if err != nil {
	// 	log.Fatalf("there was an error: %v", err())
	// } else {
	// 	fmt.Println("no err")
	// }

	// panicErr := SafeExecPanic(errorFunc)
	// if panicErr != nil {
	// 	log.Fatalf("there was an Panic: %v", panicErr())
	// } else {
	// 	fmt.Println("no panic")
	// }

	noErr := SafeExec(errorFunc)
	if noErr != nil {
		log.Fatalf("there was no error: %v", noErr())
	} else {
		fmt.Println("no err or panic")
	}
}
