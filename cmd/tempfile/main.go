package main

import "os"

func main() {
	f()
	m()
	println("done.")
}

func f() {
	f, err := os.CreateTemp(".", "b*c")
	if err != nil {
		panic(err)
	}
	f.Close()
}

func m() {
	p, err := os.MkdirTemp(".", "12*")
	if err != nil {
		panic(err)
	}
	println(p)
}
