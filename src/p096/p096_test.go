package p096

import (
	"testing"
)

func testCase(in, want int, t *testing.T) {
	got := numTrees(in)
	if got != want {
		t.Errorf("Case %v: expected %v, got %v", in, want, got)
	}
}

func TestNumTrees(t *testing.T) {
	in0 := 3
	want0 := 5
	testCase(in0, want0, t)
}

func BenchmarkNumTrees(b *testing.B) {
	in := 19
	want := 1767263190

	for i := 0; i < b.N; i++ {
		got := numTrees(in)
		if got != want {
			b.Errorf("Benchmark case %v: expected %v, got %v", in, want, got)
		}
	}
}
