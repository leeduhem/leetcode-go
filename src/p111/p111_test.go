package p111

import (
	"testing"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func testCase(in *TreeNode, want int, t *testing.T) {
	got := minDepth(in)
	if got != want {
		t.Errorf("Case %v: expected %v, got %v", in, want, got)
	}
}

func TestMinDepth(t *testing.T) {
	in0 := &TreeNode{3, &TreeNode{2, &TreeNode{1, nil, nil}, nil},
		&TreeNode{4, nil, nil}}
	want0 := 2
	testCase(in0, want0, t)

	in1 := &TreeNode{1, &TreeNode{2, nil, nil}, nil}
	want1 := 2
	testCase(in1, want1, t)
}
