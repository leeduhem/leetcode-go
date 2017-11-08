package p101

import (
	"testing"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func testCase(in *TreeNode, want bool, t *testing.T) {
	got := isSymmetric(in)
	if got != want {
		t.Errorf("Case %v: expected %v, got %v", in, want, got)
	}
}

func TestIsSymmetric(t *testing.T) {
	in0 := &TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}},
		&TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{3, nil, nil}}}
	want0 := true
	testCase(in0, want0, t)

	in1 := &TreeNode{1, &TreeNode{2, nil, &TreeNode{3, nil, nil}},
		&TreeNode{2, nil, &TreeNode{3, nil, nil}}}
	want1 := false
	testCase(in1, want1, t)

	in2 := &TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, nil},
		&TreeNode{3, &TreeNode{2, nil, nil}, nil}}
	want2 := false
	testCase(in2, want2, t)
}
