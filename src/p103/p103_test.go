package p103

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
	got := zigzagLevelOrder(in)
	if ! Is2DIntsEqual(got, want) {
		t.Errorf("Case %v: expected %v, got %v", in, want, got)
	}
}

func TestZigzagLevelOrder(t *testing.T) {
	in0 := &TreeNode{3, &TreeNode{9, nil, nil},
		&TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	want0 := [][]int {
		{3},
		{20,9},
		{15,7},
	}
	testCase(in0, want0, t)

	in1 := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}},
		&TreeNode{3, nil, nil}}
	want1 := [][]int {
		{1},
		{3,2},
		{4,5},
	}
	testCase(in1, want1, t)

	in2 := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, nil},
		&TreeNode{3, nil, &TreeNode{5, nil, nil}}}
	want2 := [][]int {
		{1},
		{3,2},
		{4,5},
	}
	testCase(in2, want2, t)
}
