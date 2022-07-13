package main

import "fmt"

type GT interface {
	int | string
}

type Hi[T GT] interface {
	Hello(t T)
}

// 泛型的使用:通过类型来使用泛型,而不是通过方法的泛型参数来使用
type obj[T GT] struct {
}

func (o *obj[T]) Hi(t T) {
	fmt.Println(t)
}

func main() {
	fmt.Println("vim-go")

	var o obj[int]
	o.Hi(123)

	var o2 obj[string]
	o2.Hi("abc")
}
