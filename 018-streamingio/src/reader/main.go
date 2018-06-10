package main

import (
	"fmt"
	"io"
	"strings"
)

func NewAlphabetReader(r io.Reader) *AlphabetReader {
	return &AlphabetReader{reader: r}
}

type AlphabetReader struct {
	reader io.Reader
}

func (ar *AlphabetReader) Read(p []byte) (int, error) {
	n, _ := ar.reader.Read(p)
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		if char := alpha(p[i]); char != 0 {
			buf[i] = char
		}
	}
	copy(p, buf) // very important
	return 0, nil
}

func alpha(b byte) byte {
	if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
		return b
	}
	return 0
}

func main() {
	r := NewAlphabetReader(strings.NewReader("12123123jungle is a fucnkknfla3241345k13'5'r dgisD"))

	p := make([]byte, 256)
	r.Read(p)
	fmt.Println(string(p))
}
