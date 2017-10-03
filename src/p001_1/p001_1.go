package p001_1

func twoSum(nums []int, target int) []int {
	var m = make(map[int]int, len(nums))

	for i,v := range nums {
		var res = target - v
		var j,found = m[res]
		if found {
			return []int {j, i}
		}

		m[v] = i
	}

	panic("invalid inputs")
}
