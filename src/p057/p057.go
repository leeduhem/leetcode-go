package p057

import "sort"

func merge(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return []Interval {}
	}

	sort.Slice(intervals, func (i,j int) bool { return intervals[i].Start < intervals[j].Start })

	ins := []Interval { intervals[0] }
	for i := 1; i < len(intervals); i++ {
		last := &(ins[len(ins)-1])
		cur  := intervals[i]

		if cur.Start > last.End {
			ins = append(ins, cur)
		} else if cur.End > last.End {
			last.End = cur.End
		}
	}

	return ins
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	return merge(append(intervals, newInterval))
}
