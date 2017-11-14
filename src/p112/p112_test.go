package p112

import (
	"testing"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func testCase(in0 *TreeNode, in1 int, want bool, t *testing.T) {
	got := hasPathSum(in0, in1)
	if got != want {
		t.Errorf("Case %v, %v: expected %v, got %v", in0, in1, want, got)
	}
}

func TestHasPathSum(t *testing.T) {
	in00 := &TreeNode{5, &TreeNode{4, &TreeNode{11, &TreeNode{7, nil, nil}, &TreeNode{2, nil, nil}}, nil},
		&TreeNode{8, &TreeNode{13, nil, nil}, &TreeNode{4, nil, &TreeNode{1, nil, nil}}}}
	in01 := 22
	want0 := true
	testCase(in00, in01, want0, t)

	in10 := (*TreeNode)(nil)
	in11 := 0
	want1 := false
	testCase(in10, in11, want1, t)

	in20 := &TreeNode{1, &TreeNode{2, nil, nil}, nil}
	in21 := 1
	want2 := false
	testCase(in20, in21, want2, t)
}
