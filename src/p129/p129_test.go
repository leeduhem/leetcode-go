package p129

import (
	"testing"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func testCase(in *TreeNode, want int, t *testing.T) {
	got := sumNumbers(in)
	if got != want {
		t.Errorf("Case %v: expected %v, got %v", in, want, got)
	}
}

func TestSumNumbers(t *testing.T) {
	in0 := &TreeNode{1, &TreeNode{ Val : 2 }, &TreeNode{ Val : 3 }}
	want0 := 25
	testCase(in0, want0, t)
}
