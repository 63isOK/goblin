package main

import (
	"fmt"
	"time"
)

func main() {
	var i int = 10

	p(i)

	println("=============")
	newType()
}

func p(i interface{}) {
	fmt.Println(i.(int))
}

type T time.Time

func (t *T) P() {
	println("123")
}

func newType() {
	var t T
	t.P()
	// t不再拥有time.Time的方法集
}
