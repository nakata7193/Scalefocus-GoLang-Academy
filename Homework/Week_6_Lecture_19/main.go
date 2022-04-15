package main

import (
	"io"
	"os"
	"strings"
)

type ReverseStringReader struct {
	r io.Reader
}

func newReverseStringReader(r io.Reader) *ReverseStringReader {
	return &ReverseStringReader{r: r}
}

func (r *ReverseStringReader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	if err != nil {
		return
	}
	for i := 0; i < n/2; i++ {
		p[i], p[n-i-1] = p[n-i-1], p[i]
	}
	return
}

func main() {
	r := newReverseStringReader(strings.NewReader("Hello, world!"))
	io.Copy(os.Stdout, r)
}
