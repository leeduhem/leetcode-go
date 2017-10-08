package p031

import "testing"

func testCase(in0, out0 []int, t *testing.T) {
	oin0 := make([]int, len(in0))
	copy(oin0, in0)

	nextPermutation(in0)

	if len(in0) != len(out0) {
		t.Fail()
	}

	if ! t.Failed() {
		for i,v := range in0 {
			if v != out0[i] {
				t.Fail()
			}
		}
	}

	if t.Failed() {
		t.Errorf("Case %v: expected %v, got %v", oin0, out0, in0)
	}
}

func TestNextPermutation(t *testing.T) {
	in10  := []int {1, 2, 3}
	out10 := []int {1, 3, 2}
	testCase(in10, out10, t)

	in20  := []int {3, 2, 1}
	out20 := []int {1, 2, 3}
	testCase(in20, out20, t)

	in30  := []int {1, 1, 5}
	out30 := []int {1, 5, 1}
	testCase(in30, out30, t)

	in40  := []int {1, 2}
	out40 := []int {2, 1}
	testCase(in40, out40, t)

	in50  := []int {2, 3, 1}
	out50 := []int {3, 1, 2}
	testCase(in50, out50, t)
}
