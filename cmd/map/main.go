package main

func main() {
	a := make(map[int]string)
	a[1] = "a"
	a[2] = "b"
	a[3] = "c"

	for v := range a {
		println(v)
	}
}
