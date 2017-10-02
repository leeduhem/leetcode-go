package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i,vi := range nums {
		for j,vj := range (nums[i+1:]) {
			if (vi + vj == target) {
				return []int {i,i+j+1}
			}
		}
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
