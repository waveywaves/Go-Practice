package main

import (
	"bytes"
	"io"
	"os"
)

func main() {

	proverbs := new(bytes.Buffer) // zeroed storage
	proverbs.WriteString("This is Awesome !")
	proverbs.WriteString("Awesome this is !")
	proverbs.WriteString("To believe is the first step to succeed !")

	var f *os.File

	f, _ = os.Create("./thisisafilereally")
	defer f.Close()
	io.Copy(f, proverbs)
}
