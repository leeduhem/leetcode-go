package utils

import (
	"sort"
	"strconv"
	"strings"
)

func Ints2Key(ns []int) string {
	sort.Ints(ns)

	ss := []string {}
	for _,n := range ns {
		ss = append(ss, strconv.Itoa(n))
	}

	return strings.Join(ss, ",")
}

func CmpMatrix(in0, in1 [][]int) bool {
	if len(in0) != len(in1) {
		return false
	}

	m := map[string]int {}
	for i,v := range in0 {
		k := Ints2Key(v)
		m[k] = i
	}

	for _,v := range in1 {
		k := Ints2Key(v)
		if _,ok := m[k]; !ok {
			return false
		}
	}

	return true
}
