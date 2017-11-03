package p091

import (
	"testing"
)

func testCase(in string, want int, t *testing.T) {
	got := numDecodings(in)
	if got != want {
		t.Errorf("Case %v: expected %v, got %v", in, want, got)
	}
}

func TestNumDecodings(t *testing.T) {
	in0 := "12"
	want0 := 2
	testCase(in0, want0, t)

	in1 := "123"
	want1 := 3
	testCase(in1, want1, t)

	in2 := "0"
	want2 := 0
	testCase(in2, want2, t)

	in3 := "27"
	want3 := 1
	testCase(in3, want3, t)

	in4 := "012"
	want4 := 0
	testCase(in4, want4, t)

	in5 := ""
	want5 := 0
	testCase(in5, want5, t)

	in6 := "10"
	want6 := 1
	testCase(in6, want6, t)
}

func BenchmarkNumDecodings(b *testing.B) {
	in := "9371597631128776948387197132267188677349946742344217846154932859125134924241649584251978418763151253"
	want := 3981312

	for i := 0; i < b.N; i++ {
		got := numDecodings(in)
		if got != want {
			b.Errorf("Benchmark #0: expected %v, got %v", want, got)
		}
	}
}
