package p167_1

func twoSum(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return []int {-1, -1}
	}

	left, right := 0, len(numbers)-1
	for left < right {
		n := numbers[right] + numbers[left]
		if n == target {
			return []int {left + 1, right + 1}
		}

		if n > target {
			right--
		} else {
			left++
		}
	}

	return []int {-1, -1}
}
