package p001

func twoSum(nums []int, target int) []int {
	for i,vi := range nums {
		for j,vj := range (nums[i+1:]) {
			if (vi + vj == target) {
				return []int {i,j+i+1}
			}
		}
	}

	panic("invalid inputs")
}
