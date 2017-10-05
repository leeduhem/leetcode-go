package p018

import(
	"sort"
	"strings"
	"strconv"
)

func ints2key(ns []int) string {
	sort.Ints(ns)

	ss := make([]string, len(ns))

	for i,n := range ns {
		ss[i] = strconv.Itoa(n)
	}

	return strings.Join(ss, ",")
}

func fourSum(nums []int, target int) [][]int {
	res := [][]int {}
	resm := map[string]int {}

	for i,vi := range nums {
		for j,vj := range nums[i+1:] {
			for k,vk := range nums[j+i+2:] {
				for l,vl := range nums[k+j+i+3:] {
					sum := vi + vj + vk + vl
					if sum == target {
						ints := []int {vi, vj, vk, vl}
						key := ints2key(ints)
						if _,ok := resm[key]; !ok {
							res = append(res, []int{vi, vj, vk, vl})
							resm[key] = l
						}
					}
				}
			}
		}
	}

	return res
}
