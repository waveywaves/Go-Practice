package main

import (
	"bytes"
	"fmt"
)

func main() {
	proverbs := []string{
		"Channels orchestrate, mutexes serialize\n",
		"Errors are values\n",
		"Don't panic\n",
	}
	var buf bytes.Buffer
	b := make([]byte, 256)
	for i, v := range proverbs {
		fmt.Printf("Writing :- \n\t %v : %v \n", i, v)
		buf.Write([]byte(v))
	}
	fmt.Printf("byteformat : %v\n", buf)
	n, _ := buf.Read(b)
	fmt.Printf("Values read from buffer: %v \nstringformat : %v", n, string(b))
}
