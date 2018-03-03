package p148

func merge(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	var prev *ListNode

	for prev = head; l1 != nil && l2 != nil; prev = prev.Next {
		if l1.Val < l2.Val {
			prev.Next = l1;
			l1 = l1.Next
		} else {
			prev.Next = l2;
			l2 = l2.Next
		}
	}

	if l1 != nil {
		prev.Next = l1
	}
	if l2 != nil {
		prev.Next = l2
	}

	return head.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// Step 1: cut the list to two halves
	var prev, slow *ListNode = nil, head

	for fast := head; fast != nil && fast.Next != nil; fast = fast.Next.Next {
		prev, slow = slow, slow.Next
	}

	prev.Next = nil

	// Step 2: Sort each half
	l1 := sortList(head)
	l2 := sortList(slow)

	// Step 3: Merge l1 and l2
	return merge(l1, l2)
}
