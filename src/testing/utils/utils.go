package utils

import (
	"fmt"
	"strconv"
	"strings"
)

type IntMatrix [][]int

func (m IntMatrix) String() string {
	str := "["
	for i,r := range m {
		if i > 0 { str += " " }
		if len(r) > 0 {
			str += fmt.Sprintf("[%v", r[0])
			for j := 1; j < len(r); j++ {
				str += fmt.Sprintf(" %v", r[j])
			}
			str += "]"
		} else {
			str += "[]"
		}

		if i != len(m)-1 {
			str += "\n"
		}
	}
	str += "]"

	return str
}

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

func DeepCopy2DMatrix(m [][]int) [][]int {
	m1 := make([][]int, len(m))
	for i,r := range m {
		m1[i] = make([]int, len(r))
		copy(m1[i], r)
	}
	return m1
}

func Is2DIntMatrixEqual(m, m1 [][]int) bool {
	if len(m) != len(m1) {
		return false
	}

	for i,r := range m {
		for j,c := range r {
			if c != m1[i][j] {
				return false
			}
		}
	}

	return true
}

func Is2DStrsEqualWithoutOrder(in0, in1 [][]string) bool {
	if len(in0) != len(in1) {
		return false
	}

	m := map[string]int {}
	for i,ss := range in0 {
		k := strings.Join(ss, ",")
		m[k] = i
	}

	for _,ss := range in1 {
		k := strings.Join(ss, ",")
		delete(m, k)
	}

	return len(m) == 0
}
