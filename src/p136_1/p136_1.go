package p136_1

func singleNumber(nums []int) int {
	// A XOR A = 0 and the XOR operator is commutative
	res := 0
	for _, n := range nums {
		res ^= n
	}
	return res
}
