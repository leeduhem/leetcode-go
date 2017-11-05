package p095

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

func isTreesEqual(trees1, trees2 []*TreeNode) bool {
	if trees1 == nil && trees2 == nil {
		return true
	}

	if trees1 != nil && trees2 != nil {
	outerLoop:
		for _, t1 := range trees1 {
			for _, t2 := range trees2 {
				if isTreeEqual(t1, t2) {
					continue outerLoop
				}
			}
			return false
		}

		return true
	}

	return false
}

func testCase(in int, want []*TreeNode, t *testing.T) {
	got := generateTrees(in)
	if ! isTreesEqual(got, want) {
		t.Errorf("Case %v: expected \n%v\n, got \n%v\n", in, want, got)
	}
}

func TestGenerateTrees(t *testing.T) {
	in0 := 3
	want0 := []*TreeNode{
		&TreeNode{1, nil, &TreeNode{3, &TreeNode{2, nil, nil}, nil}},
		&TreeNode{3, &TreeNode{2, &TreeNode{1, nil, nil}, nil}, nil},
		&TreeNode{3, &TreeNode{1, nil, &TreeNode{2, nil, nil}}, nil},
		&TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}},
		&TreeNode{1, nil, &TreeNode{2, nil, &TreeNode{3, nil, nil}}},
	}
	testCase(in0, want0, t)
}
