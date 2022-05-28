package main

import "reflect"

// ValueOf传入的都是拷贝

func main() {
	s := make([]int, 0)
	v := reflect.ValueOf(&s)

	println(v.CanSet())

	println(reflect.ValueOf(make(chan int)).CanSet())
}
