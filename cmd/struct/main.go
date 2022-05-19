package main

import "reflect"

type A struct{}

func (a *A) Hello() {
	println("A")
}

func (a *A) HelloReal() {
	ref := reflect.ValueOf(a)
	method := ref.MethodByName("Hello")
	if method.IsValid() {
		method.Call(make([]reflect.Value, 0))
	} else {
		// 错误处理
	}
}

type B struct {
	A
}

func (b *B) Hello() {
	println("B")
}

func (b *B) HelloReal() {
	b.Hello()
}

func main() {
	var b B
	b.Hello()

	b.A.HelloReal()
}

// 基类可以持有一个接口,子类实现接口,这样就可以在基类调用子类的方法
