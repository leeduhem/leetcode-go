package p066

func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits)-1; i >= 0; i-- {
		d1 := digits[i] + carry
		if d1 >= 10 {
			carry = 1
			d1 -= 10
		} else {
			carry = 0
		}
		digits[i] = d1
	}

	if carry > 0 {
		return append([]int{carry}, digits...)
	}
	return digits
}
