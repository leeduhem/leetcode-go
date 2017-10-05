package p016_1

import "testing"

func testCase(ns []int, t0, r0 int, t *testing.T) {
	r1 := threeSumClosest(ns, t0)
	if r0 != r1 {
		t.Errorf("Case %v: expected %v, got %v", ns, r0, r1)
	}
}

func TestThreeSumClosest(t *testing.T) {
	ns1 := []int {-1, 2, 1, -4}
	r10 := 2
	testCase(ns1, 1, r10, t)

	ns2 := []int {1, 1, -1}
	r20 := 1
	testCase(ns2, 2, r20, t)
}
