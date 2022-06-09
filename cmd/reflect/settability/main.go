package main

import "reflect"

// ValueOf传入的都是拷贝

type home struct {
	a int
	B int
}

func main() {
	s := make([]int, 0)
	v := reflect.ValueOf(&s)

	println(v.CanSet())

	println(reflect.ValueOf(make(chan int)).CanSet())

	var h home
	v = reflect.ValueOf(&h).Elem()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		println(f.Type().Name(), v.CanSet())
	}
}
