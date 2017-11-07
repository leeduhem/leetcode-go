package p099_1

import (
	"testing"
	"strings"
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// Poor man's tree printing function...
func (r *TreeNode) String() string {
	var printNode func (level int, node *TreeNode) string;
	printNode = func (level int, node *TreeNode) string {
		str := strings.Repeat("-", level)
		if node == nil {
			str += "<nil>\n"
		} else {
			str += fmt.Sprintf("%v\n", node.Val)
			str += printNode(level+1, node.Left)
			str += printNode(level+1, node.Right)
		}
		return str
	}
	return printNode(1, r) + "\n"
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

func deepCopyTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	return &TreeNode{ root.Val, deepCopyTree(root.Left), deepCopyTree(root.Right) }
}

func testCase(in, want *TreeNode, t *testing.T) {
	oin := deepCopyTree(in)
	recoverTree(in)
	if ! isTreeEqual(in, want) {
		t.Errorf("Case\n%v: expected\n%v, got\n%v", oin, want, in)
	}
}

func TestRecoverTree(t *testing.T) {
	in0   := &TreeNode{10, &TreeNode{5, nil, nil}, &TreeNode{15, &TreeNode{6, nil, nil}, &TreeNode{20, nil, nil}}}
	want0 := &TreeNode{6, &TreeNode{5, nil, nil}, &TreeNode{15, &TreeNode{10, nil, nil}, &TreeNode{20, nil, nil}}}
	testCase(in0, want0, t)

	in1   := &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{1, nil, nil}}
	want1 := &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}
	testCase(in1, want1, t)

	in2   := &TreeNode{1, nil, &TreeNode{3, nil, &TreeNode{2, nil, nil}}}
	want2 := &TreeNode{1, nil, &TreeNode{2, nil, &TreeNode{3, nil, nil}}}
	testCase(in2, want2, t)
}
