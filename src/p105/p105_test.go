package p105

import (
	"testing"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isTreeEqual(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}

	if t1 != nil && t2 != nil {
		return t1.Val == t2.Val &&
			isTreeEqual(t1.Left, t2.Left) &&
			isTreeEqual(t1.Right, t2.Right)
	}

	return false
}

func testCase(in0, in1 []int, want *TreeNode, t *testing.T) {
	got := buildTree(in0, in1)
	if ! isTreeEqual(got, want) {
		t.Errorf("Case %v, %v: expected %v, got %v", in0, in1, want, got)
	}
}

func TestBuildTree(t *testing.T) {
	in00 := []int {1, 2, 3}
	in01 := []int {2, 1, 3}
	want0 := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	testCase(in00, in01, want0, t)

	in10 := []int {1, 2}
	in11 := []int {1, 2}
	want1 := &TreeNode{1, nil, &TreeNode{2, nil, nil}}
	testCase(in10, in11, want1, t)

	in20 := []int {1,2,3}
	in21 := []int {3,2,1}
	want2 := &TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, nil}, nil}
	testCase(in20, in21, want2, t)

	in30 := []int {3,2,1,4}
	in31 := []int {1,2,3,4}
	want3 := &TreeNode{3, &TreeNode{2, &TreeNode{1, nil, nil}, nil},
		&TreeNode{4, nil, nil}}
	testCase(in30, in31, want3, t)
}
