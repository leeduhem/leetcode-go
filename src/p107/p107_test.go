package p107

import (
	"testing"
	. "testing/utils"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func testCase(in *TreeNode, want [][]int, t *testing.T) {
	got := levelOrderBottom(in)
	if ! Is2DIntsEqual(got, want) {
		t.Errorf("Case %v: expected %v, got %v", in, want, got)
	}
}

func TestLevelOrder(t *testing.T) {
	in0 := &TreeNode{3, &TreeNode{9, nil, nil},
		&TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	want0 := [][]int {
		{15,7},
		{9,20},
		{3},
	}
	testCase(in0, want0, t)
}
