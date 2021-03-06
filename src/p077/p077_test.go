package p077

import (
	"testing"
	. "testing/utils"
)

func testCase(in0, in1 int, want [][]int, t *testing.T) {
	got := combine(in0, in1)
	if ! Is2DIntsEqualWithoutOrder(want, got) {
		t.Errorf("Case %v, %v: expected %v, got %v", in0, in1, want, got)
	}
}

func TestCombine(t *testing.T) {
	in00  := 4
	in01  := 2
	want0 := [][]int {
		{2,4},
		{3,4},
		{2,3},
		{1,2},
		{1,3},
		{1,4},
	}
	testCase(in00, in01, want0, t)
}

func BenchmarkCombine(b *testing.B) {
	in0, in1 := 5, 4
	want := [][]int {
		{2,3,4,5},
		{1,3,4,5},
		{1,2,4,5},
		{1,2,3,5},
		{1,2,3,4},
	}

	for i := 0; i < b.N; i++ {
		got := combine(in0, in1)
		if ! Is2DIntsEqualWithoutOrder(want, got) {
			b.Errorf("Benchmark #0: expected %v, got %v", want, got)
		}
	}
}
