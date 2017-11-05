package p094

import (
	"testing"
	. "testing/utils"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func testCase(in *TreeNode, want []int, t *testing.T) {
	got := inorderTraversal(in)
	if ! IsIntsEqual(got, want) {
		t.Errorf("Case %v: expected %v, got %v", in, want, got)
	}
}

func TestInorderTraversal(t *testing.T) {
	in0 := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}
	want0 := []int {1,3,2}
	testCase(in0, want0, t)
}
