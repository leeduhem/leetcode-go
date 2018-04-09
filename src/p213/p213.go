package p213

func intMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func lineRob(nums []int) int {
	var pre, cur int
	for _, n := range nums {
		tmp := intMax(pre+n, cur)
		pre, cur = cur, tmp
	}
	return cur
}

func rob(nums []int) int {
	switch ln := len(nums); ln {
	case 0:
		return 0
	case 1:
		return nums[0]
	default:
		money1 := lineRob(nums[:ln-1])
		money2 := lineRob(nums[1:])
		return intMax(money1, money2)
	}
}
