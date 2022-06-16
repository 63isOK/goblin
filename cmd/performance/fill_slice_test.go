package performance_test

import (
	"testing"

	. "github.com/63isOK/goblin/cmd/performance"
)

func TestWithAppend(t *testing.T) {
	got := WithAppend()

	if len(got) != 1 {
		t.Fatalf("预计长度大于1,实际%d", len(got))
	}
}

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
