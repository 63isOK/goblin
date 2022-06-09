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
	println("=============")
	newAliasType()
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

type B time.Time

func (b *B) Day() int {
	return 123
}

func newAliasType() {
	t := time.Now()
	println(t.Day())

	var b B
	println(b.Day())
	// 类型别名也失去了方法级
	// b.GoString()
}
