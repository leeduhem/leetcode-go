package p052

import "testing"

func testCase(in0, out0 int, t *testing.T) {
	out1 := totalNQueens(in0)
	if out0 != out1 {
		t.Errorf("Case %v: expected %v, got %v", in0, out0, out1)
	}
}

func TestSolveNQueens(t *testing.T) {
	in1  := 4
	out1 := 2
	testCase(in1, out1, t)

	in2  := 8
	out2 := 92
	testCase(in2, out2, t)
}
