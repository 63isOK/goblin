package main

func main() {
	p()
	pif()
	pif2()
	pf()
	pfor()
}

func p() {
	a, b := 1, 2
	d := &a
	a, c := 3, 4

	println(a, b, c, *d)
}

func pif() {
	println("if")
	a, b := 1, 2
	d := &a
	if a, c := 3, 4; c > 3 {
		println(a, b, c, *d)
	}

	println(a, b, *d)
}

func pif2() {
	println("if2")
	a, b := 1, 2
	d := &a
	if c := 4; c > 3 {
		a, c := 3, 4
		println(a, b, c, *d)
	}

	println(a, b, *d)
}

func pf() {
	println("pf")
	var i int = 3
	p := &i
	pi(p)

	println(*p)
}

func pi(i *int) {
	println("pi")
	p := &i
	println(*i, **p)
	a := 2
	b, i := a, &a
	println(b, *i, **p)
}

func pfor() {
	println("pfor")
	i := 0
	p := &i
	for i, j := 1, 2; j < 3; j++ {

		println(i, j, *p)
	}

	println(*p)
}

func preturn() (i *int) {
	println("preturn")
	p := &i
	var a int = 1
	i, j := &a, 2
	println(*i, j, *p)

	return
}
