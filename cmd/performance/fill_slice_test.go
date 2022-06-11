package performance_test

import (
	"testing"

	. "github.com/63isOK/goblin/cmd/performance"
)

func BenchmarkWithAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WithAppend()
	}
}

func BenchmarkWithAppendAndLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WithAppendAndLength()
	}
}

func BenchmarkWithAssign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WithAssign()
	}
}
