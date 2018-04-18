// https://leetcode.com/problems/add-two-numbers/description/
package p002

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := &ListNode{0, nil}
	carry := 0

	for tail := sum; ; tail = tail.Next {
		s := carry
		if l1 == nil && l2 == nil {
			if carry != 0 {
				tail.Next = &ListNode{1, nil}
			}
			break
		}

		if l1 != nil {
			s += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			s += l2.Val
			l2 = l2.Next
		}

		if s > 9 {
			s -= 10
			carry = 1
		} else {
			carry = 0
		}

		tail.Next = &ListNode{s, nil}
	}

	return sum.Next
}
