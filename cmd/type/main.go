package main

import "fmt"

func main() {
	var i int = 10

	p(i)
}

func p(i interface{}) {
	fmt.Println(i.(string))
}
