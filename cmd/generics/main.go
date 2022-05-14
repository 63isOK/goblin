package main

import "golang.org/x/exp/constraints"

func main() {
	mi := map[string]int64{
		"one": 1,
		"two": 2,
	}

	println(sumInts(mi))

	mf := map[string]float64{
		"one": 1.1,
		"two": 2.2,
	}

	println(sumFloats(mf))

	println("=========================")

	println(sumIntsOrFloats(mi))
	println(sumIntsOrFloats(mf))
}

func sumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}

	return s
}

func sumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}

	return s
}

func sumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}

	return s
}

func GMin[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}
