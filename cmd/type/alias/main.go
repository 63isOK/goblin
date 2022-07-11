package main

type A struct {
	Name string
}

func (a *A) P() {
	println(a.Name)
}

func (a *A) p() {
	println("hi", a.Name)
}

// 类型别名不会丢失方法集
type B = A

func main() {
	a := A{Name: "tom"}
	a.P()
	a.p()

	println()

	b := B{Name: "jim"}
	b.P()
	b.p()
}
