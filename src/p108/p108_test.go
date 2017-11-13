package p108

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

func testCase(in []int, want *TreeNode, t *testing.T) {
	got := sortedArrayToBST(in)
	if ! isTreeEqual(got, want) {
		t.Errorf("Case %v: Expected %v, got %v", in, want, got)
	}
}

func TestSortedArrayToBST(t *testing.T) {
	in0 := []int {1, 2, 3, 4}
	want0 := &TreeNode{3, &TreeNode{2, &TreeNode{1, nil, nil}, nil},
		&TreeNode{4, nil, nil}}
	testCase(in0, want0, t)
}
