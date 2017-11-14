package p113

import (
	"testing"
	. "testing/utils"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func is2DIntsEqualWithouFirstLevelOrder(nss1, nss2 [][]int) bool {
	if len(nss1) != len(nss2) {
		return false
	}

outerLoop:
	for _, ns1 := range nss1 {
		for _, ns2 := range nss2 {
			if IsIntsEqual(ns1, ns2) {
				continue outerLoop
			}
		}
		return false
	}

	return true
}

func testCase(in0 *TreeNode, in1 int, want [][]int, t *testing.T) {
	got := pathSum(in0, in1)
	if ! is2DIntsEqualWithouFirstLevelOrder(got, want) {
		t.Errorf("Case %v, %v: expected %v, got %v", in0, in1, want, got)
	}
}

func TestHasPathSum(t *testing.T) {
	in00 := &TreeNode{5, &TreeNode{4, &TreeNode{11, &TreeNode{7, nil, nil}, &TreeNode{2, nil, nil}}, nil},
		&TreeNode{8, &TreeNode{13, nil, nil}, &TreeNode{4, &TreeNode{5, nil, nil}, &TreeNode{1, nil, nil}}}}
	in01 := 22
	want0 := [][]int {
		{5,4,11,2},
		{5,8,4,5},
	}
	testCase(in00, in01, want0, t)
}
