package main

import (
	"fmt"

	"github.com/63isOK/goblin/cmd/type/alias/pkg/raw"
	nt "github.com/63isOK/goblin/cmd/type/alias/pkg/type1"
)

func main() {
	o := raw.New()
	p(o)
}

// 别名是可以进行传递的,实际上她们都是等价的
func p(o nt.OT) {
	fmt.Printf("%v", o)
}
