package p117

import (
	"testing"
)

type TreeLinkNode struct {
	Val int
	Left, Right, Next *TreeLinkNode
}

func deepCopyTreeLinkNode(root *TreeLinkNode) *TreeLinkNode {
	if root == nil {
		return nil
	}

	return &TreeLinkNode{ root.Val, deepCopyTreeLinkNode(root.Left), deepCopyTreeLinkNode(root.Right), root.Next }
}

func isLinkedListEqual(list1, list2 *TreeLinkNode) bool {
	if list1 == nil && list2 == nil {
		return true
	}

	if list1 != nil && list2 != nil {
		return list1.Val == list2.Val &&
			isLinkedListEqual(list1.Next, list2.Next)
	}

	return false
}

func isTreeLinkNodeEqual(root1, root2 *TreeLinkNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 != nil && root2 != nil {
		return root1.Val == root2.Val &&
			isTreeLinkNodeEqual(root1.Left, root2.Left) &&
			isTreeLinkNodeEqual(root1.Right, root2.Right) &&
			isLinkedListEqual(root1.Next, root2.Next)
	}

	return false
}

func testCase(in, want *TreeLinkNode, t *testing.T) {
	oin := deepCopyTreeLinkNode(in)
	connect(in)
	if ! isTreeLinkNodeEqual(in, want) {
		t.Errorf("Case %v: expected %v, got %v", oin, want, in)
	}
}

func TestConnect(t *testing.T) {
	in0 := &TreeLinkNode{ Val: 1, Left: &TreeLinkNode{ Val: 2, Left: &TreeLinkNode{ Val: 4 }, Right: &TreeLinkNode{ Val: 5 }},
		Right: &TreeLinkNode{ Val: 3, Right: &TreeLinkNode{ Val: 7 }}}
	want0 := &TreeLinkNode{ Val: 1, Left: &TreeLinkNode{ Val: 2, Left: &TreeLinkNode{ Val: 4 }, Right: &TreeLinkNode{ Val: 5 }},
		Right: &TreeLinkNode{ Val: 3, Right: &TreeLinkNode{ Val: 7 }}}
	// level 1
	want0.Left.Next = want0.Right
	// level 2
	want0.Left.Left.Next = want0.Left.Right
	want0.Left.Right.Next = want0.Right.Right

	testCase(in0, want0, t)
}
