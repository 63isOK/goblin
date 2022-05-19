package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushBack(123)
	l.PushBack("hello")

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%#v, %T\n", e.Value, e.Value)
	}
}
