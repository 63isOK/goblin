package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("vim-go")
	var i int = 3
	f := c(i)
	a := f()
	fmt.Println(reflect.TypeOf(a).String())
	if _, ok := a.(*int); !ok {
		panic("not *int")
	}
}

// 深拷贝
func c(i any) func() any {
	return func() any {
		return reflect.New(reflect.TypeOf(i)).Interface()
	}
}
