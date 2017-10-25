package p072

import (
	"testing"
)

func testCase(in0, in1 string, want int, t *testing.T) {
	got := minDistance(in0, in1)
	if got != want {
		t.Errorf("Case %v %v: expected %v, got %v", in0, in1, want, got)
	}
}

func TestMinDistance(t *testing.T) {
	cases := []struct {
		in0, in1 string
		want     int
	} {
		{"abc", "ac", 1},
		{   "",   "", 0},
		{  "a",  "b", 1},
	}

	for _,c := range cases {
		testCase(c.in0, c.in1, c.want, t)
	}
}

func BenchmarkMinDistance(b *testing.B) {
	in0  := "dinitrophenylhydrazine"
	in1  := "benzalphenylhydrazone"
	want := 7

	for i := 0; i < b.N; i++ {
		got := minDistance(in0, in1)
		if got != want {
			b.Errorf("Benchmark #0: expected %v, got %v", want, got)
		}
	}
}
