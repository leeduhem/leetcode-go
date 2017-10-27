package p078

func deepCopy(m [][]int) [][]int {
	m1 := make([][]int, len(m))
	for i,r := range m {
		m1[i] = make([]int, len(r))
		copy(m1[i], r)
	}
	return m1
}

func subsets(nums []int) [][]int {
	ss := [][]int { {} }
	if len(nums) == 0 {
		return ss
	}

	ss1 := subsets(nums[1:])
	ss = append(ss[1:], ss1...) // Ignore extra empty subset

	ss2 := deepCopy(ss1)
	for i, s := range ss2 {
		ss2[i] = append(s, nums[0])
	}
	ss = append(ss, ss2...)

	return ss
}
