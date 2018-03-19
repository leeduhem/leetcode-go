package p198

func intMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func rob(nums []int) int {
	var pre, cur int
	for _, n := range nums {
		tmp := intMax(pre + n, cur)
		pre, cur = cur, tmp
	}
	return cur
}
