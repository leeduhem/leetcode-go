package p124

import (
	"testing"
)

func testCase(in *TreeNode, want int, t *testing.T) {
	got := maxPathSum(in)
	if got != want {
		t.Errorf("Case %v: expected %v, got %v", in, want, got)
	}
}

func TestMaxPathSum(t *testing.T) {
	in0 := &TreeNode{1, &TreeNode{ Val : 2 }, &TreeNode{ Val : 3 }}
	want0 := 6
	testCase(in0, want0, t)
}
