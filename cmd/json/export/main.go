package main

import (
	"encoding/json"
	"fmt"
)

// json库只能看到main.ABC的暴露字段
type ABC struct {
	a int
	b int
	C string
}

func main() {
	o := ABC{
		a: 1, b: 2, C: "abc",
	}
	s, err := json.Marshal(&o)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(s))
}
