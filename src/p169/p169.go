package p169

func majorityElement(nums []int) int {
	m := map[int]int {}

	for _, n := range nums {
		m[n] += 1
	}

	for k, v := range m {
		if v > len(nums) / 2 {
			return k
		}
	}

	panic("Invalid input")
}
