package utils

import (
	"fmt"
	"sort"
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
	sort.Ints(ns)
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

func DeepCopy2DIntMatrix(m [][]int) [][]int {
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

func Is2DIntsEqual(nss1, nss2 [][]int) bool {
	if len(nss1) != len(nss2) {
		return false
	}

	for i, ns1 := range nss1 {
		if len(ns1) != len(nss2[i]) {
			return false
		}

		for j, n1 := range ns1 {
			if n1 != nss2[i][j] {
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

func DeepCopyInts(in []int) []int {
	in1 := make([]int, len(in))
	copy(in1, in)
	return in1
}

func IsIntsEqual(ns, ns1 []int) bool {
	if len(ns) != len(ns1) {
		return false
	}

	for i,n := range ns {
		if n != ns1[i] {
			return false
		}
	}

	return true
}

func IsStrsEqual(ss, ss1 []string) bool {
	if len(ss) != len(ss1) {
		return false
	}

	for i,s := range ss {
		if s != ss1[i] {
			return false
		}
	}

	return true
}

func IsStrsEqualWithoutOrder(ss, ss1 []string) bool {
	if len(ss) != len(ss1) {
		return false
	}

	m := map[string]int {}
	for i := 0; i < len(ss); i++ {
		m[ss[i]]++
		m[ss1[i]]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}

func IsOneOf(ns []int, nss [][]int) bool {
	for _, ns1 := range nss {
		if IsIntsEqual(ns, ns1) {
			return true
		}
	}
	return false
}
