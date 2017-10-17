package p041

import "testing"

func testCase(in0 []int, out0 int, t *testing.T) {
	out1 := firstMissingPositive(in0)
	if out0 != out1 {
		t.Errorf("Case %v: expected %v, got %v", in0, out0, out1)
	}
}

func TestFirstMissingPositive(t *testing.T) {
	in10  := []int {1, 2, 0}
	out10 := 3
	testCase(in10, out10, t)

	in20  := []int {3, 4, -1, 1}
	out20 := 2
	testCase(in20, out20, t)

	in30  := []int {2}
	out30 := 1
	testCase(in30, out30, t)

	in40  := []int {1000, -1}
	out40 := 1
	testCase(in40, out40, t)

	in50  := []int {0,2,2,4,0,1,0,1,3}
	out50 := 5
	testCase(in50, out50, t)
}
