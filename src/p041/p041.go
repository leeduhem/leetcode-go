package p041

import "container/list"

type closedInterval struct {
	low, high int
}

func insertEndpoint(p int, l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		v := e.Value.(closedInterval)

		if v.low <= p && p <= v.high {
			return
		}

		if p + 1 == v.low {
			v.low = p; e.Value = v
			if e1 := e.Prev(); e1 != nil {
				v1 := e1.Value.(closedInterval)
				if v1.high + 1 == v.low {
					v2 := closedInterval{ v1.low, v.high }
					l.InsertBefore(v2, e1)
					l.Remove(e1)
					l.Remove(e)
				}
			}
			return
		}

		if v.high + 1 == p {
			v.high = p; e.Value = v
			if e1 := e.Next(); e1 != nil {
				v1 := e1.Value.(closedInterval)
				if v.high + 1 == v1.low {
					v2 := closedInterval{ v.low, v1.high }
					l.InsertBefore(v2, e1)
					l.Remove(e1)
					l.Remove(e)
				}
			}
			return
		}

		if p < v.low {
			v1 := closedInterval{ p, p }
			l.InsertBefore(v1, e)
			return
		}
	}

	l.PushBack(closedInterval{p, p})
}

func firstMissingPositive(nums []int) int {
	l := list.New()

	for _,n := range nums {
		insertEndpoint(n, l)
	}

	for e := l.Front(); e != nil; e = e.Next() {
		v := e.Value.(closedInterval)

		if v.high <= 0 {
			if e1 := e.Next(); e1 != nil {
				v1 := e1.Value.(closedInterval)
				if v1.high <= 0 {
					continue
				}

				if v1.low > 1 {
					return 1
				}
				return v1.high + 1
			}
			return 1
		}

		if v.low > 1 {
			return 1
		}
		return v.high + 1
	}

	return 1
}
