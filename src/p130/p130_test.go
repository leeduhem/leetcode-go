package p130

import (
	"testing"
)

func deepCopyBoard(b [][]byte) [][]byte {
	b1 := make([][]byte, len(b))
	for i, r := range b {
		b1[i] = make([]byte, len(r))
		copy(b1[i], r)
	}
	return b1
}

func isBoardEqual(b1, b2 [][]byte) bool {
	if len(b1) != len(b2) {
		return false
	}

	for i1, r1 := range b1 {
		if len(b2[i1]) != len(r1) {
			return false
		}

		for j1, c1 := range r1 {
			if c1 != b2[i1][j1] {
				return false
			}
		}
	}

	return true
}

func testCase(in, want [][]byte, t *testing.T) {
	oin := deepCopyBoard(in)
	solve(in)
	if ! isBoardEqual(in, want) {
		t.Errorf("Case %v: expected %v, got %v", oin, want, in)
	}
}

func TestSolve(t *testing.T) {
	in0 := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	want0 := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	testCase(in0, want0, t)
}
