package main

import (
	"fmt"
	"io"
)

type obj struct {
	i io.Reader
}

func main() {
	p()
	println("=================")
	cmp()
}

func p() {
	var s obj

	fmt.Printf("%+v\n", s)
}

func cmp() {
	var i interface{}
	var age int = 322
	var address string = "abc"

	i = age
	i = nil
	if i == 12 {
		println("i=12")
	} else if i == 32 {
		println("i=32")
	} else if i == "abc" {
		println("i=abc")
	} else if i == nil {
		println("i=nil")
	}

	i = address
	if i == "abc" {
		println("i=abc")
	}
}
