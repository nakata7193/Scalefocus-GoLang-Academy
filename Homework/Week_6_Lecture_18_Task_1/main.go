package main

import (
	"io"
	"log"
	"os"
)

type ReverseStringReader struct {
	read   io.Reader
	output string
}

func NewReverseStringReader(input string) *ReverseStringReader {

	r := []rune(input) //inverting the string
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	newReverseStringReader := &ReverseStringReader{output: string(r)}
	return newReverseStringReader
}

func main() {

	file, err := os.Open("dat.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	var result []byte
	buf := make([]byte, 1024)
	for {
		n, err := ReverseStringReader.read.Read()
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if err == io.EOF {
			break
		}
		result = append(result, buf[:n]...)
	}

}
