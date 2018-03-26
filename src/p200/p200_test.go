package p200

import (
	"testing"
)

func deepCopy(grid [][]byte) [][]byte {
	grid1 := make([][]byte, len(grid))

	for i, row := range grid {
		grid1[i] = make([]byte, len(row))
		copy(grid1[i], row)
	}

	return grid1
}

func TestNumIslands(t *testing.T) {
	testCase := func(in [][]byte, want int) {
		in0 := deepCopy(in)
		got := numIslands(in)
		if got != want {
			t.Errorf("Case %v: expected %v, got %v", in0, want, got)
		}
	}

	in0 := [][]byte {
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	testCase(in0, 1)

	in1 := [][]byte {
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	testCase(in1, 3)
}
