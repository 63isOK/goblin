package main

import (
	"fmt"
	"io"
)

type obj struct {
	i io.Reader
}

func main() {
	var s obj

	fmt.Printf("%+v", s)
}
