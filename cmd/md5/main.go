package main

import (
	"crypto/md5"
	"fmt"
)

func encodeName2Key(name string) string {
	return fmt.Sprintf("%X", (md5.Sum([]byte(name))))
}

func main() {
	println(encodeName2Key("abc"))
}
