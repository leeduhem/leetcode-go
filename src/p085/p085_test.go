package p085

import (
	"testing"
)

import "fmt"
type Bins []byte

func (b Bins) String() string {
	str := "["
	for _, c := range b {
		str += fmt.Sprintf(" %c", c)
	}
	str += "]"

	return str
}

func testCase(in [][]byte, want int, t *testing.T) {
	got := maximalRectangle(in)
	if got != want {
		t.Errorf("Case %v: expected %v, got %v", in, want, got)
	}
}

func TestMaximalRectangle(t *testing.T) {
	in0 := [][]byte {
		{'1','0','1','0','0'},
		{'1','0','1','1','1'},
		{'1','1','1','1','1'},
		{'1','0','0','1','0'},
	}
	want0 := 6
	testCase(in0, want0, t)

	in1 := [][]byte {
		{'1','0','1','1','1'},
		{'0','1','0','1','0'},
		{'1','1','0','1','1'},
		{'1','1','0','1','1'},
		{'0','1','1','1','1'},
	}
	want1 := 6
	testCase(in1, want1, t)
}
