package p149

import (
	"testing"
)

type Point struct {
	X, Y int
}

func testCase(in []Point, want int, t *testing.T) {
	got := maxPoints(in)
	if got != want {
		t.Errorf("Case %v: expected %v, got %v", in, want, got)
	}
}

func TestMaxPoints(t *testing.T) {
	in0 := []Point {}
	want0 := 0
	testCase(in0, want0, t)

	in1 := []Point {Point{0,0}, Point{1,1}, Point{1,-1}}
	want1 := 2
	testCase(in1, want1, t)
}
