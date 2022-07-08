package main

import "fmt"

type A struct {
	Hi    string
	Hello string
}
type B struct {
	Hi    string
	Hello string
}
type C struct {
	Hi    string
	Hello string
}

type o struct {
	A
	B
	C
	Hi    string
	Hello string
}

func main() {
	var obj o
	obj.A.Hi = "1"
	obj.A.Hello = "2"
	obj.B.Hi = "3"
	obj.B.Hello = "4"
	obj.C.Hi = "5"
	obj.C.Hello = "6"

	fmt.Printf("%+v", obj)

	fmt.Println(obj.Hi, obj.Hello)
}
