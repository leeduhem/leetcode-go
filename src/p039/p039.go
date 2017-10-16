package p039

import (
	"sort"
	"strconv"
	"strings"
)

func ints2key(ns []int) string {
	sort.Ints(ns)

	ss := []string {}
	for _,n := range ns {
		ss = append(ss, strconv.Itoa(n))
	}

	return strings.Join(ss, ",")
}

var mem = map[string]bool {}

func checkPartialCombination(candidates []int, target int, partial []int) [][]int {
	plen := len(partial)
	results := [][]int {}

	for _,n := range candidates {
		if target < n {
			continue
		}

		// Appending to `partial' directly will change the `partial'
		// in calling function.
		p := make([]int, plen + 1)
		copy(p, partial)
		p[plen] = n

		key := ints2key(p)
		if _,ok := mem[key]; ok {
			continue
		}

		if target == n {
			results = append(results, p)
			mem[key] = true
		} else {
			results1 := checkPartialCombination(candidates, target-n, p)
			results = append(results, results1...)
		}
	}

	return results
}

func combinationSum(candidates []int, target int) [][]int {
	results := [][]int {}

	// clear memory
	mem = map[string]bool {}

	for _,n := range candidates {
		if target < n {
			continue
		}

		// candidates does not include duplicates
		p := []int {n}
		if target == n {
			results = append(results, p)
		} else {
			results1 := checkPartialCombination(candidates, target-n, p)
			results = append(results, results1...)
		}
	}

	return results
}
