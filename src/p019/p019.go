package p019

type ListNode struct {
    Val int
    Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pres := [](*ListNode) {nil}

	for pre := head; pre != nil; pre = pre.Next {
		pres = append(pres, pre)
	}

	len0 := len(pres)
	ind  := len0 - n - 1
	if pres[ind] == nil {
		return pres[1].Next
	}
	if ind == len0 - 2 {
		pres[ind].Next = nil
		return pres[1]
	}

	pres[ind].Next = pres[ind+2]
	return pres[1]
}
