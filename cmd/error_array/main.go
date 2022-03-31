package main

import (
	"errors"
	"fmt"
)

func main() {
	errs := []error{
		errors.New("123"),
		errors.New("1"),
		errors.New("a"),
		errors.New("b"),
	}

	fmt.Printf("打印错误数组:(%s)", errs)
}
