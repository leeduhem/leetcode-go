package p036

import "testing"

func testCase(in0 [][]byte, out0 bool, t *testing.T) {
	out1 := isValidSudoku(in0)
	if out0 != out1 {
		t.Errorf("Case %v: expected %v, got %v", in0, out0, out1)
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
	out10 := true
	testCase(in10, out10, t)

	in20  := [][]byte {
		[]byte(".87654321"),
		[]byte("2........"),
		[]byte("3........"),
		[]byte("4........"),
		[]byte("5........"),
		[]byte("6........"),
		[]byte("7........"),
		[]byte("8........"),
		[]byte("9........"),
	}
	out20 := true
	testCase(in20, out20, t)
}
