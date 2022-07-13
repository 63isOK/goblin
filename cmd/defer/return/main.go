package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Printf("%v", p())
}

// 在defer中只能通过具名返回值来修改函数返回参数
func p() (err error) {
	defer func() {
		if err != nil {
			err = errors.New("1234")
		}
	}()

	// err = errors.New("abc")
	// return err

	// 直接return 会更新具名返回值,然后通过defer来修改
	return errors.New("abc")
}
