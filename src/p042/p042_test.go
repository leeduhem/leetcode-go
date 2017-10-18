package p042

import "testing"

func testCase(in0 []int, out0 int, t *testing.T) {
	out1 := trap(in0)
	if out0 != out1 {
		t.Errorf("Case %v: expected %v, got %v", in0, out0, out1)
	}
}

func TestTrap(t *testing.T) {
	in10  := []int {0,1,0,2,1,0,1,3,2,1,2,1}
	out10 := 6
	testCase(in10, out10, t)

	in20  := []int {}
	out20 := 0
	testCase(in20, out20, t)

	in30  := []int {2, 0, 2}
	out30 := 2
	testCase(in30, out30, t)

	in40  := []int {5,2,1,2,1,5}
	out40 := 14
	testCase(in40, out40, t)
}
