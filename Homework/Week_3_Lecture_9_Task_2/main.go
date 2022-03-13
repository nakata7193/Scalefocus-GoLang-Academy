package main

type Action func() error

type Pipe struct{}

func NewPipe() *Pipe {
	return &Pipe{}
}

func SafeExec(a Action) (f Action) {

	return f
}

// func (p *Pipe) Do() error {
// 	var err error
// 	return fmt.Errorf("Safe exec: %w",err)
// }

func main() {
	pipe := &Pipe{}

	err := pipe.SafeExec()

	if err != nil {
		return err
	}
}
