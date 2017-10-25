package p072

import (
	"container/heap"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }


func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func min(ns []int) int {
	// Oh, this is a little bit of overkill.
	heap.Init((*IntHeap)(&ns))
	return ns[0]
}

var mem map[[2]int]int

// https://en.wikipedia.org/wiki/Edit_distance
func minDistance(word1, word2 string) int {
	// clean memory
	mem = map[[2]int]int {}

	if len(word1) == 0 { return len(word2) }
	if len(word2) == 0 { return len(word1) }

	var distance func(int, int) int;
	distance = func (i,j int) int {
		key := [2]int{i,j}
		if d,ok := mem[key]; ok {
			return d
		}

		if j == -1 { return i + 1 }
		if i == -1 { return j + 1 }

		if word1[j] == word2[i] {
			d := distance(i-1,j-1)
			mem[key] = d
			return d
		}

		d := 1 + min([]int {
			distance(i-1,  j),
			distance(i,  j-1),
			distance(i-1,j-1),
		})
		mem[key] = d
		return d
	}
	return distance(len(word2)-1, len(word1)-1)
}
