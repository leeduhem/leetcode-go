package p034_1

func searchRangeShort(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int {-1, -1}
	}

	lo, hi := -1, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			lo = i
			break
		}
	}

	if lo < 0 {
		return []int {-1, -1}
	}

	for j := len(nums)-1; lo <= j; j-- {
		if nums[j] == target {
			hi = j
			break
		}
	}

	return []int {lo, hi}
}


func searchRange(nums []int, target int) []int {
	nlen := len(nums)

	if nlen < 16 {
		return searchRangeShort(nums, target)
	}

	h := (nlen - 1) / 2
	if nums[h] == target {
		if nums[h+1] > target {
			return searchRange(nums[:h+1], target)
		} else {
			// nums is sorted in ascending order
			// nums[h+1] == target
			inds1 := searchRange(nums[:h+1], target)
			inds2 := searchRange(nums[h+1:], target)
			return []int {inds1[0], h+1 + inds2[1]}
		}
	} else if nums[h] < target {
		inds1 := searchRange(nums[h:], target)
		if inds1[0] >= 0 {
			return []int {h + inds1[0], h + inds1[1]}
		}
		return []int {-1, -1}
	} else {
		// nums[h] > target
		return searchRange(nums[:h], target)
	}
}
