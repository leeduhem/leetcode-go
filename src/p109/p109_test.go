package p109

import (
	"testing"
)

func testCase(in *ListNode, want *TreeNode, t *testing.T) {
	oin := deepCopyList(in)
	got := sortedListToBST(in)
	if ! isTreeEqual(got, want) {
		t.Errorf("Case %v: expected %v, got %v", oin, want, got)
	}
}

func TestSortedListToBST(t *testing.T) {
	in0 := ints2List([]int {1, 2, 3, 4})
	want0 := &TreeNode{3, &TreeNode{2, &TreeNode{1, nil, nil}, nil},
		&TreeNode{4, nil, nil}}
	testCase(in0, want0, t)
}
