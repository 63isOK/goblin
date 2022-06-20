package main

func main() {
	var b bo

	b.p()
}

type ao int

func (a ao) p() {
	println("a")
}

type bo struct {
	ao
}

// 之类是可以调用父类的方法的
func (b bo) p() {
	b.ao.p()
	println("b")
}
