package p002

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sum = &ListNode{ 0, nil }
	var end = sum
	var carry = 0

	for {
		var s = carry
		if (l1 == nil && l2 == nil) {
			if (carry != 0) {
				end.Next = &ListNode { 1, nil }
				end = end.Next
			}
			break
		}

		if (l1 != nil) {
			s += l1.Val
			l1 = l1.Next
		}
		if (l2 != nil) {
			s += l2.Val
			l2 = l2.Next
		}

		if (s > 9) {
			s -= 10
			carry = 1
		} else {
			carry = 0
		}

		end.Next = &ListNode{ s, nil }
		end = end.Next
	}

	return sum.Next
}
