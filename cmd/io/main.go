package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	var r io.Reader = strings.NewReader("some io.Reader stream to be read\n")

	r = io.TeeReader(r, os.Stdout)

	// Everything read from r will be copied to stdout.
	s, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}
	println(string(s))
}
