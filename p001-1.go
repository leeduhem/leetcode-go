package main

import "fmt"

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

func main() {
	var numbers = []int {3,2,4}
	var target = 6

	var indins = twoSum(numbers, target)
	for idx, val := range indins {
		fmt.Println(idx, ":", val)
	}
}
