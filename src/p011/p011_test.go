package p011

import "testing"

func testCase(hs []int, a0 int, t *testing.T) {
	a1 := maxArea(hs)
	if a1 != a0 {
		t.Errorf("Case %v: expected %v, got %v", hs, a0, a1)
	}
}

func TestMaxArea(t *testing.T) {
	testCase([]int {1,1}, 1, t)
}
