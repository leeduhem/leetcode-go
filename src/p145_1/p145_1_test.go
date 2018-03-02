package p145_1

import (
	"testing"
	. "testing/utils"
)

func testCase(in *TreeNode, want []int, t *testing.T) {
	got := postorderTraversal(in)
	if ! IsIntsEqual(got, want) {
		t.Errorf("Case %v: expected %v, got %v", in, want, got)
	}
}

func TestPostorderTraversal(t *testing.T) {
	in0 := &TreeNode{ Val: 1, Right: &TreeNode{ Left: &TreeNode{ Val: 3 }, Val: 2 }}
	want0 := []int { 3, 2, 1 }
	testCase(in0, want0, t)
}
