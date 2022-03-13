package main

type Action func() error

func SafeExec(a Action) Action {
	return nil
}

func main() {
	
}