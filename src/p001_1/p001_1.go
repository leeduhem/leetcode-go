// https://leetcode.com/problems/two-sum/description/
package p001_1

func twoSum(nums []int, target int) []int {
	var m = make(map[int]int, len(nums))

	for i, v := range nums {
		res := target - v
		if j, ok := m[res]; ok {
			return []int{j, i}
		}
		m[v] = i
	}

	panic("invalid inputs")
}
