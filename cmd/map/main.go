package main

func main() {
	a := make(map[int]string)
	a[1] = "a"
	a[2] = "b"
	a[3] = "c"

	for v := range a {
		println(v)
	}

	var i string
	i = a[4]
	println(i)

	p()
}

func p() {
	a := make(map[int]int)
	println(a[12])
	println(len(a))
}
