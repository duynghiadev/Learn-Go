package main

import (
	"golang.org/x/tour/reader"
)

// MyReader emits an infinite stream of the ASCII character 'A'.
type MyReader struct{}

// Read fills the byte slice with the ASCII character 'A'.
func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
