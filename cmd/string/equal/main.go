package main

import (
	"fmt"
	"strings"
)

// 大小写不敏感的比较
func main() {
	fmt.Printf("%v\t", strings.EqualFold(`abc`, "ABc"))
	fmt.Printf("%v\t", strings.EqualFold(`abc`, "aBc"))
	fmt.Printf("%v\t", strings.EqualFold("abc", "ABc"))
	fmt.Printf("%v\t", strings.EqualFold(string('a'), "A"))
}
