package performance

var src = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 1, 1, 1, 1, 11, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 11, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 11, 1, 1, 1, 1, 1, 1, 11, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 11, 1, 2, 2, 2, 2, 22, 2, 2, 2, 22, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 11, 1}

// var src = []int{"string", "1", 1, 1, 1, 1, 11, 1, 1, 1, "abc"}

func WithAppend() []int {
	dst := make([]int, 0)
	for _, v := range src {
		dst = append(dst, v+1)
	}

	return dst
}

func WithAppendAndLength() []int {
	dst := make([]int, 0, len(src))
	for _, v := range src {
		dst = append(dst, v+1)
	}

	return dst
}

func WithAssign() []int {
	dst := make([]int, len(src))
	for i, v := range src {
		dst[i] = v + 1
	}

	return dst
}
