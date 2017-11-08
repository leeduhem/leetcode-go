package p100

import (
	"testing"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func testCase(in0, in1 *TreeNode, want bool, t *testing.T) {
	got := isSameTree(in0, in1)
	if got != want {
		t.Errorf("Case %v, %v: expected %v, got %v", in0, in1, want, got)
	}
}

func TestIsSameTree(t *testing.T) {
	in00 := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	in01 := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	want0 := true
	testCase(in00, in01, want0, t)

	in10 := &TreeNode{1, &TreeNode{2, nil, nil}, nil}
	in11 := &TreeNode{1, nil, &TreeNode{2, nil, nil}}
	want1 := false
	testCase(in10, in11, want1, t)

	in20 := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{1, nil, nil}}
	in21 := &TreeNode{1, &TreeNode{1, nil, nil}, &TreeNode{2, nil, nil}}
	want2 := false
	testCase(in20, in21, want2, t)
}
