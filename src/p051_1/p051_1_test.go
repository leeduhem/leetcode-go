package p051_1

import "testing"
import . "testing/utils"

func testCase(in0 int, out0 [][]string, t *testing.T) {
	out1 := solveNQueens(in0)
	if ! Is2DStrsEqualWithoutOrder(out0, out1) {
		t.Errorf("Case %v: expected %v, got %v", in0, out0, out1)
	}
}

func TestSolveNQueens(t *testing.T) {
	in1  := 4
	out1 := [][]string {
		{".Q..",
		 "...Q",
		 "Q...",
		 "..Q.",},

		{"..Q.",
		 "Q...",
		 "...Q",
		 ".Q..",},
	}
	testCase(in1, out1, t)

	in2  := 8
	out2 := [][]string {
		{"Q.......", "....Q...", ".......Q", ".....Q..", "..Q.....", "......Q.", ".Q......", "...Q....", },
		{"Q.......", ".....Q..", ".......Q", "..Q.....", "......Q.", "...Q....", ".Q......", "....Q...", },
		{"Q.......", "......Q.", "...Q....", ".....Q..", ".......Q", ".Q......", "....Q...", "..Q.....", },
		{"Q.......", "......Q.", "....Q...", ".......Q", ".Q......", "...Q....", ".....Q..", "..Q.....", },
		{".Q......", "...Q....", ".....Q..", ".......Q", "..Q.....", "Q.......", "......Q.", "....Q...", },
		{".Q......", "....Q...", "......Q.", "Q.......", "..Q.....", ".......Q", ".....Q..", "...Q....", },
		{".Q......", "....Q...", "......Q.", "...Q....", "Q.......", ".......Q", ".....Q..", "..Q.....", },
		{".Q......", ".....Q..", "Q.......", "......Q.", "...Q....", ".......Q", "..Q.....", "....Q...", },
		{".Q......", ".....Q..", ".......Q", "..Q.....", "Q.......", "...Q....", "......Q.", "....Q...", },
		{".Q......", "......Q.", "..Q.....", ".....Q..", ".......Q", "....Q...", "Q.......", "...Q....", },
		{".Q......", "......Q.", "....Q...", ".......Q", "Q.......", "...Q....", ".....Q..", "..Q.....", },
		{".Q......", ".......Q", ".....Q..", "Q.......", "..Q.....", "....Q...", "......Q.", "...Q....", },
		{"..Q.....", "Q.......", "......Q.", "....Q...", ".......Q", ".Q......", "...Q....", ".....Q..", },
		{"..Q.....", "....Q...", ".Q......", ".......Q", "Q.......", "......Q.", "...Q....", ".....Q..", },
		{"..Q.....", "....Q...", ".Q......", ".......Q", ".....Q..", "...Q....", "......Q.", "Q.......", },
		{"..Q.....", "....Q...", "......Q.", "Q.......", "...Q....", ".Q......", ".......Q", ".....Q..", },
		{"..Q.....", "....Q...", ".......Q", "...Q....", "Q.......", "......Q.", ".Q......", ".....Q..", },
		{"..Q.....", ".....Q..", ".Q......", "....Q...", ".......Q", "Q.......", "......Q.", "...Q....", },
		{"..Q.....", ".....Q..", ".Q......", "......Q.", "Q.......", "...Q....", ".......Q", "....Q...", },
		{"..Q.....", ".....Q..", ".Q......", "......Q.", "....Q...", "Q.......", ".......Q", "...Q....", },
		{"..Q.....", ".....Q..", "...Q....", "Q.......", ".......Q", "....Q...", "......Q.", ".Q......", },
		{"..Q.....", ".....Q..", "...Q....", ".Q......", ".......Q", "....Q...", "......Q.", "Q.......", },
		{"..Q.....", ".....Q..", ".......Q", "Q.......", "...Q....", "......Q.", "....Q...", ".Q......", },
		{"..Q.....", ".....Q..", ".......Q", "Q.......", "....Q...", "......Q.", ".Q......", "...Q....", },
		{"..Q.....", ".....Q..", ".......Q", ".Q......", "...Q....", "Q.......", "......Q.", "....Q...", },
		{"..Q.....", "......Q.", ".Q......", ".......Q", "....Q...", "Q.......", "...Q....", ".....Q..", },
		{"..Q.....", "......Q.", ".Q......", ".......Q", ".....Q..", "...Q....", "Q.......", "....Q...", },
		{"..Q.....", ".......Q", "...Q....", "......Q.", "Q.......", ".....Q..", ".Q......", "....Q...", },
		{"...Q....", "Q.......", "....Q...", ".......Q", ".Q......", "......Q.", "..Q.....", ".....Q..", },
		{"...Q....", "Q.......", "....Q...", ".......Q", ".....Q..", "..Q.....", "......Q.", ".Q......", },
		{"...Q....", ".Q......", "....Q...", ".......Q", ".....Q..", "Q.......", "..Q.....", "......Q.", },
		{"...Q....", ".Q......", "......Q.", "..Q.....", ".....Q..", ".......Q", "Q.......", "....Q...", },
		{"...Q....", ".Q......", "......Q.", "..Q.....", ".....Q..", ".......Q", "....Q...", "Q.......", },
		{"...Q....", ".Q......", "......Q.", "....Q...", "Q.......", ".......Q", ".....Q..", "..Q.....", },
		{"...Q....", ".Q......", ".......Q", "....Q...", "......Q.", "Q.......", "..Q.....", ".....Q..", },
		{"...Q....", ".Q......", ".......Q", ".....Q..", "Q.......", "..Q.....", "....Q...", "......Q.", },
		{"...Q....", ".....Q..", "Q.......", "....Q...", ".Q......", ".......Q", "..Q.....", "......Q.", },
		{"...Q....", ".....Q..", ".......Q", ".Q......", "......Q.", "Q.......", "..Q.....", "....Q...", },
		{"...Q....", ".....Q..", ".......Q", "..Q.....", "Q.......", "......Q.", "....Q...", ".Q......", },
		{"...Q....", "......Q.", "Q.......", ".......Q", "....Q...", ".Q......", ".....Q..", "..Q.....", },
		{"...Q....", "......Q.", "..Q.....", ".......Q", ".Q......", "....Q...", "Q.......", ".....Q..", },
		{"...Q....", "......Q.", "....Q...", ".Q......", ".....Q..", "Q.......", "..Q.....", ".......Q", },
		{"...Q....", "......Q.", "....Q...", "..Q.....", "Q.......", ".....Q..", ".......Q", ".Q......", },
		{"...Q....", ".......Q", "Q.......", "..Q.....", ".....Q..", ".Q......", "......Q.", "....Q...", },
		{"...Q....", ".......Q", "Q.......", "....Q...", "......Q.", ".Q......", ".....Q..", "..Q.....", },
		{"...Q....", ".......Q", "....Q...", "..Q.....", "Q.......", "......Q.", ".Q......", ".....Q..", },
		{"....Q...", "Q.......", "...Q....", ".....Q..", ".......Q", ".Q......", "......Q.", "..Q.....", },
		{"....Q...", "Q.......", ".......Q", "...Q....", ".Q......", "......Q.", "..Q.....", ".....Q..", },
		{"....Q...", "Q.......", ".......Q", ".....Q..", "..Q.....", "......Q.", ".Q......", "...Q....", },
		{"....Q...", ".Q......", "...Q....", ".....Q..", ".......Q", "..Q.....", "Q.......", "......Q.", },
		{"....Q...", ".Q......", "...Q....", "......Q.", "..Q.....", ".......Q", ".....Q..", "Q.......", },
		{"....Q...", ".Q......", ".....Q..", "Q.......", "......Q.", "...Q....", ".......Q", "..Q.....", },
		{"....Q...", ".Q......", ".......Q", "Q.......", "...Q....", "......Q.", "..Q.....", ".....Q..", },
		{"....Q...", "..Q.....", "Q.......", ".....Q..", ".......Q", ".Q......", "...Q....", "......Q.", },
		{"....Q...", "..Q.....", "Q.......", "......Q.", ".Q......", ".......Q", ".....Q..", "...Q....", },
		{"....Q...", "..Q.....", ".......Q", "...Q....", "......Q.", "Q.......", ".....Q..", ".Q......", },
		{"....Q...", "......Q.", "Q.......", "..Q.....", ".......Q", ".....Q..", "...Q....", ".Q......", },
		{"....Q...", "......Q.", "Q.......", "...Q....", ".Q......", ".......Q", ".....Q..", "..Q.....", },
		{"....Q...", "......Q.", ".Q......", "...Q....", ".......Q", "Q.......", "..Q.....", ".....Q..", },
		{"....Q...", "......Q.", ".Q......", ".....Q..", "..Q.....", "Q.......", "...Q....", ".......Q", },
		{"....Q...", "......Q.", ".Q......", ".....Q..", "..Q.....", "Q.......", ".......Q", "...Q....", },
		{"....Q...", "......Q.", "...Q....", "Q.......", "..Q.....", ".......Q", ".....Q..", ".Q......", },
		{"....Q...", ".......Q", "...Q....", "Q.......", "..Q.....", ".....Q..", ".Q......", "......Q.", },
		{"....Q...", ".......Q", "...Q....", "Q.......", "......Q.", ".Q......", ".....Q..", "..Q.....", },
		{".....Q..", "Q.......", "....Q...", ".Q......", ".......Q", "..Q.....", "......Q.", "...Q....", },
		{".....Q..", ".Q......", "......Q.", "Q.......", "..Q.....", "....Q...", ".......Q", "...Q....", },
		{".....Q..", ".Q......", "......Q.", "Q.......", "...Q....", ".......Q", "....Q...", "..Q.....", },
		{".....Q..", "..Q.....", "Q.......", "......Q.", "....Q...", ".......Q", ".Q......", "...Q....", },
		{".....Q..", "..Q.....", "Q.......", ".......Q", "...Q....", ".Q......", "......Q.", "....Q...", },
		{".....Q..", "..Q.....", "Q.......", ".......Q", "....Q...", ".Q......", "...Q....", "......Q.", },
		{".....Q..", "..Q.....", "....Q...", "......Q.", "Q.......", "...Q....", ".Q......", ".......Q", },
		{".....Q..", "..Q.....", "....Q...", ".......Q", "Q.......", "...Q....", ".Q......", "......Q.", },
		{".....Q..", "..Q.....", "......Q.", ".Q......", "...Q....", ".......Q", "Q.......", "....Q...", },
		{".....Q..", "..Q.....", "......Q.", ".Q......", ".......Q", "....Q...", "Q.......", "...Q....", },
		{".....Q..", "..Q.....", "......Q.", "...Q....", "Q.......", ".......Q", ".Q......", "....Q...", },
		{".....Q..", "...Q....", "Q.......", "....Q...", ".......Q", ".Q......", "......Q.", "..Q.....", },
		{".....Q..", "...Q....", ".Q......", ".......Q", "....Q...", "......Q.", "Q.......", "..Q.....", },
		{".....Q..", "...Q....", "......Q.", "Q.......", "..Q.....", "....Q...", ".Q......", ".......Q", },
		{".....Q..", "...Q....", "......Q.", "Q.......", ".......Q", ".Q......", "....Q...", "..Q.....", },
		{".....Q..", ".......Q", ".Q......", "...Q....", "Q.......", "......Q.", "....Q...", "..Q.....", },
		{"......Q.", "Q.......", "..Q.....", ".......Q", ".....Q..", "...Q....", ".Q......", "....Q...", },
		{"......Q.", ".Q......", "...Q....", "Q.......", ".......Q", "....Q...", "..Q.....", ".....Q..", },
		{"......Q.", ".Q......", ".....Q..", "..Q.....", "Q.......", "...Q....", ".......Q", "....Q...", },
		{"......Q.", "..Q.....", "Q.......", ".....Q..", ".......Q", "....Q...", ".Q......", "...Q....", },
		{"......Q.", "..Q.....", ".......Q", ".Q......", "....Q...", "Q.......", ".....Q..", "...Q....", },
		{"......Q.", "...Q....", ".Q......", "....Q...", ".......Q", "Q.......", "..Q.....", ".....Q..", },
		{"......Q.", "...Q....", ".Q......", ".......Q", ".....Q..", "Q.......", "..Q.....", "....Q...", },
		{"......Q.", "....Q...", "..Q.....", "Q.......", ".....Q..", ".......Q", ".Q......", "...Q....", },
		{".......Q", ".Q......", "...Q....", "Q.......", "......Q.", "....Q...", "..Q.....", ".....Q..", },
		{".......Q", ".Q......", "....Q...", "..Q.....", "Q.......", "......Q.", "...Q....", ".....Q..", },
		{".......Q", "..Q.....", "Q.......", ".....Q..", ".Q......", "....Q...", "......Q.", "...Q....", },
		{".......Q", "...Q....", "Q.......", "..Q.....", ".....Q..", ".Q......", "......Q.", "....Q...", },
	}
	testCase(in2, out2, t)
}
