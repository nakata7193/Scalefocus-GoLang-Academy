package main

import (
	"io"
	"strings"
	"testing"
)

func TestReverseStringBuilder(t *testing.T) {
	r := strings.NewReader("Hello, world!")
	e := &ReverseStringReader{r: r}
	b := &strings.Builder{}
	io.Copy(b, e)
	if b.String() != "!dlrow ,olleH" {
		t.Errorf("Expected '!dlrow ,olleH', got '%s'", b.String())
	}
}
