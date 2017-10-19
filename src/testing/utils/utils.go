package utils

import (
	"strconv"
	"strings"
)

func Ints2Key(ns []int) string {
	ss := []string {}
	for _,n := range ns {
		ss = append(ss, strconv.Itoa(n))
	}

	return strings.Join(ss, ",")
}

func Is2DIntsEqualWithoutOrder(in0, in1 [][]int) bool {
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
		delete(m, k)
	}

	return len(m) == 0
}
