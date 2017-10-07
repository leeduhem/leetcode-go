package p023

import "container/heap"

type ListNode struct {
    Val int
    Next *ListNode
}

type NodeHeap []*ListNode

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(n interface{}) {
	*h = append(*h, n.(*ListNode))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	l := len(old)
	n := old[l-1]
	*h = old[:l-1]
	return n
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := make(NodeHeap, 0)
	heap.Init(&h)

	for _,l := range lists {
		for n := l; n != nil; n = n.Next {
			heap.Push(&h, n)
		}
	}

	head := &ListNode{0, nil}
	tail := head

	for h.Len() > 0 {
		n := heap.Pop(&h)
		tail.Next = n.(*ListNode)
		tail = tail.Next
	}
	tail.Next = nil

	return head.Next
}
