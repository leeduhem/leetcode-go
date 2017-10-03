package p002_1

type ListNode struct {
	Val int
	Next *ListNode
}

func number2Digit(n *ListNode) <-chan int {
	d := make(chan int)
	go func() {
		for ; n != nil; n = n.Next {
			d <- n.Val
		}
		close(d)
	}()

	return d
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := &ListNode{ 0, nil }
	end := sum
	carry := 0

	ds1 := number2Digit(l1)
	ds2 := number2Digit(l2)

	for {
		d1,ok1 := <-ds1
		d2,ok2 := <-ds2

		if (!ok1 && !ok2) {
			break
		}

		s := d1 + d2 + carry
		if (s > 9) {
			s -= 10
			carry = 1
		} else {
			carry = 0
		}

		end.Next = &ListNode{ s, nil }
		end = end.Next
	}

	if (carry != 0) {
		end.Next = &ListNode{ carry, nil }
		end = end.Next
	}

	return sum.Next
}
