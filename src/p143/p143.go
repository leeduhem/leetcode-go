package p143

func reorderList(head *ListNode)  {
	ps := make([](*ListNode), 0)

	for p := head; p != nil; p = p.Next {
		ps = append(ps, p)
	}

	for i, j := 0, len(ps)-1; i <= j; i, j = i+1, j-1 {
		if i == j {
			ps[j].Next = nil
			break
		}

		ps[i].Next = ps[j]
		if i+1 < j {
			ps[j].Next = ps[i+1]
		}

		if i+1 == j {
			ps[j].Next = nil
		}
	}
}
