package main

import (
	"reflect"
)

func main() {
	var i int32
	i = '1'
	t := reflect.TypeOf(i)
	println("size:", t.Size(), "bits:", t.Bits())
}
