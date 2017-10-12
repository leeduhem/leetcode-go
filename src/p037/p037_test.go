package p037

import "testing"
import . "p037/sudoku"

func testCase(in0, out0 [][]byte, t *testing.T) {
	oin0 := make([][]byte, len(in0))
	for i,r := range in0 {
		oin0[i] = make([]byte, len(r))
		copy(oin0[i], r)
	}

	solveSudoku(in0)
	if ! IsValidSudoku(in0) {
		t.Fail()
	}
	if len(in0) != len(out0) {
		t.Fail()
	}
	if !t.Failed() {
	CheckRow:
		for i,r := range out0 {
			if len(in0[i]) != len(r) {
				t.Fail()
				break
			}
			for j := 0; j < len(r); j++ {
				if in0[i][j] != r[j] {
					t.Fail()
					break CheckRow
				}
			}
		}
	}
	if t.Failed() {
		t.Errorf("Case %v:\n expected %v,\n got %v", Board(oin0), Board(out0), Board(in0))
	}
}

func TestIsValidSudoku(t *testing.T) {
	in10  := [][]byte {
		[]byte("53..7...."),
		[]byte("6..195..."),
		[]byte(".98....6."),
		[]byte("8...6...3"),
		[]byte("4..8.3..1"),
		[]byte("7...2...6"),
		[]byte(".6....28."),
		[]byte("...419..5"),
		[]byte("....8..79"),
	}
	out10 := [][]byte {
		[]byte("534678912"),
		[]byte("672195348"),
		[]byte("198342567"),
		[]byte("859761423"),
		[]byte("426853791"),
		[]byte("713924856"),
		[]byte("961537284"),
		[]byte("287419635"),
		[]byte("345286179"),
	}
	testCase(in10, out10, t)

	in20  := [][]byte {
		[]byte("..9748..."),
		[]byte("7........"),
		[]byte(".2.1.9..."),
		[]byte("..7...24."),
		[]byte(".64.1.59."),
		[]byte(".98...3.."),
		[]byte("...8.3.2."),
		[]byte("........6"),
		[]byte("...2759.."),
	}
	out20 := [][]byte {
		[]byte("519748632"),
		[]byte("783652419"),
		[]byte("426139875"),
		[]byte("357986241"),
		[]byte("264317598"),
		[]byte("198524367"),
		[]byte("975863124"),
		[]byte("832491756"),
		[]byte("641275983"),
	}
	testCase(in20, out20, t)
}
