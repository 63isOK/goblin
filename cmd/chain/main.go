package main

func main() {
	println(p())

	// p的出参和d的入参在个数和类型上是匹配的,就能直接作为参数使用
	d(p())

	// 参数个数不一致,则p()被限制为只能充当1个参数
	// d3(p(), "hello")
	// d3(1, p())
}

func p() (int, string) {
	return 1, "abc"
}

func d(i int, s string) {
	println(i, s)
}

func d3(i int, s string, i2 int) {
	println(i, s, i2)
}
