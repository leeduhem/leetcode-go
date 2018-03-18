package p191

func hammingWeight(n uint32) int {
	ones := 0
	for ; n != 0; n >>= 1 {
		if (n & 1) != 0 {
			ones++
		}
	}
	return ones
}
