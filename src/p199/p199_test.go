package p199

import (
	"testing"
	. "testing/utils"
)

func TestRightSideView(t *testing.T) {
	testCase := func(in *TreeNode, want []int) {
		got := rightSideView(in)
		if ! IsIntsEqual(got, want) {
			t.Errorf("Case %v: expected %v, got %v", in, want, got)
		}
	}

	in0 := &TreeNode{
		Val: 1,
		Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}},
	}
	testCase(in0, []int{1,3,4})

	in1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{Val: 2},
	}
	testCase(in1, []int{1,2})

	in2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 3},
	}
	testCase(in2, []int{1,3,4})
}
