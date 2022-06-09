package main

type a int

func (i *a) Out() {
	println("123")
}

func (i *a) Print(f func()) {
	f()
}

func main() {
	var i a
	i.Print(i.Out)
}
