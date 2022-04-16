package main

import (
	"os"
	"strings"
)

func main() {
	p := "./assets/devpictures/iot/device/20220403/239995_20220403011410_42_027hPi8p9.jpg"

	mkdir(p)
}

func mkdir(p string) {
	// TODO: 待优化:入参支持路径和文件路径,优化后可提取值utils仓库
	os.MkdirAll(strings.Join(
		strings.Split(p, "/")[:len(strings.Split(p, "/"))-1], "/"),
		os.ModePerm)
}
