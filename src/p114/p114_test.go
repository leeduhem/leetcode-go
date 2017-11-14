package p114

import (
	"testing"
)

func testCase(in, want *TreeNode, t *testing.T) {
	oin := deepCopyTree(in)
	flatten(in)
	if ! isTreeEqual(in, want) {
		t.Errorf("Case %v: expected %v, got %v", oin, want, in)
	}
}

func TestFlatten(t *testing.T) {
	in0 := &TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}},
		&TreeNode{5, nil, &TreeNode{6, nil, nil}}}
	want0 := &TreeNode{1, nil, &TreeNode{2, nil, &TreeNode{3, nil, &TreeNode{4, nil, &TreeNode{5, nil, &TreeNode{6, nil, nil}}}}}}
	testCase(in0, want0, t)
}
