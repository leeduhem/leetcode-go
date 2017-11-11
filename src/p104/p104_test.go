package p104

import (
	"testing"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func testCase(in *TreeNode, want int, t *testing.T) {
	got := maxDepth(in)
	if got != want {
		t.Errorf("Case %v: expected %v, got %v", in, want, got)
	}
}

func TestMaxDepth(t *testing.T) {
	in0 := &TreeNode{1, &TreeNode{2, nil, nil}, nil}
	want0 := 2
	testCase(in0, want0, t)
}
