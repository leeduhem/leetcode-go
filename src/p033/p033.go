package p033

func searchInts(ns []int, x int) int {
	nlen := len(ns)
	if nlen == 0 {
		return -1
	}

	for l, r := 0, nlen-1; l <= r; {
		h := l + (r - l) / 2
		if ns[h] == x {
			return h
		}

		if x < ns[h] {
			r = h - 1
		} else {
			l = h + 1
		}
	}

	return -1
}

func search(nums []int, target int) int {
	nlen := len(nums)

	pivot := -1
	for i := 0; i + 1 < nlen; i++ {
		if nums[i] > nums[i+1] {
			pivot = i
			break
		}
	}

	if pivot < 0 {
		return searchInts(nums, target)
	}

	if ind := searchInts(nums[:pivot+1], target); ind >= 0 {
		return ind
	}

	if ind := searchInts(nums[pivot+1:], target); ind >= 0 {
		return pivot + 1 + ind
	}

	return -1
}
