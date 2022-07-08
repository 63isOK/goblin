package main

import (
	"fmt"
	"reflect"
)

func main() {
	test()
}

type hi interface {
	SayHi() string
	SayHello() string
}

type A struct {
}

func (a *A) SayHi() string {
	return "a say hi"
}

func (a *A) SayHello() string {
	return "a say hello"
}

type B struct {
	hi
}

func (b *B) SayHi() string {
	println("b say")
	return b.hi.SayHi()
}

func (b *B) SayHello() string {
	println("b say")
	return b.hi.SayHello()
}

func test() {
	println(getobj().SayHi())
	println(getobj().SayHello())
	var i hi = &A{}
	p(reflect.TypeOf(i))
}

func getobj() hi {
	var a A
	b := B{&a}
	return &b
}

func p(t reflect.Type) {
	for i := 0; i < t.NumMethod(); i++ {
		println(t.Method(i).Name)
		fmt.Printf("%+v", t.Method(i).Func)
	}
}
