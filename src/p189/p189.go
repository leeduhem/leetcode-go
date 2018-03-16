package p189

func rotate(nums []int, k int)  {
	ln := len(nums)
	k %= ln
	copy(nums, append(nums[ln-k:], nums[:ln-k]...))
}
