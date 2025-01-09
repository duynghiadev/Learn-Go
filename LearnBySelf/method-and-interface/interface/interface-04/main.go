package main

import (
	"io"
	"os"
	"strings"
)

// rot13Reader wraps an io.Reader and applies the ROT13 cipher to its output.
type rot13Reader struct {
	r io.Reader
}

// rot13 is a helper function that applies the ROT13 cipher to a single byte.
func rot13(b byte) byte {
	switch {
	case 'A' <= b && b <= 'Z':
		return 'A' + (b-'A'+13)%26
	case 'a' <= b && b <= 'z':
		return 'a' + (b-'a'+13)%26
	default:
		return b
	}
}

// Read reads from the underlying io.Reader, applies the ROT13 cipher, and writes to the provided byte slice.
func (r *rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b) // Read data from the underlying reader.
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i]) // Apply the ROT13 cipher to each byte.
	}
	return n, err // Return the number of bytes read and any error encountered.
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
