package main

import (
	"fmt"
	"strconv"

	"github.com/samber/lo"
)

func main() {
	fmt.Println("vim-go")
	is := []int{1, 2, 3}
	fmt.Printf("%v", lo.Map(is, func(x int, _ int) string {
		return strconv.Itoa(x)
	}))
}
