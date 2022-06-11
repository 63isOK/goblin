package performance

// var src = []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 1, 1, 1, 1, 11, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 11, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 11, 1, 1, 1, 1, 1, 1, 11, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 11, 1, 2, 2, 2, 2, 22, 2, 2, 2, 22, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 11, 1}
var src = []any{"string", "1", 1, 1, 1, 1, 11, 1, 1, 1, "abc"}

func WithAppend() []any {
	dst := make([]any, 0)
	for _, v := range src {
		dst = append(dst, v)
	}

	return dst
}

func WithAppendAndLength() []any {
	dst := make([]any, 0, len(src))
	for _, v := range src {
		dst = append(dst, v)
	}

	return dst
}

func WithAssign() []any {
	dst := make([]any, len(src))
	for i, v := range src {
		dst[i] = v
	}

	return dst
}
