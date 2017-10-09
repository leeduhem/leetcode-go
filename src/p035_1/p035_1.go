package p035_1

func searchInsert(nums []int, target int) int {
	nlen := len(nums)
	if nlen == 0 {
		return 0
	}

	for l, r := 0, nlen-1; l < r; {
		h := l + (r - l) / 2

		if nums[h] < target {
			if h+1 < nlen {
				if target <= nums[h+1] {
					return h+1
				}
				l = h+1
				continue
			}
			return l + h + 1
		}

		if nums[h] >= target {
			r = h
		}
	}

	if target <= nums[0] {
		return 0
	}
	return nlen
}
