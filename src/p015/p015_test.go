package p015

import (
	"testing"
	"sort"
	"strconv"
	"strings"
)

func ints2key(nums []int) string {
	ss := make([]string, len(nums))

	sort.Ints(nums)
	for i,v := range nums {
		ss[i] = strconv.Itoa(v)
	}

	return strings.Join(ss, ",")
}

func testCase(s []int, r0 [][]int, t *testing.T) {
	r1 := threeSum(s)

	if len(r1) != len(r0) {
		t.Fatalf("Case %v: expected %v, got %v", s, r0, r1)
	}

	rm0 := make(map[string][]int)
	for _,v := range r0 {
		k := ints2key(v)
		rm0[k] = v
	}

	for _,v := range r1 {
		key := ints2key(v)
		if _,err := rm0[key]; !err {
			t.Errorf("Case %v: expected %v, got %v", s, r0, r1)
		}
	}


}

func TestThreeSum(t *testing.T) {
	s1  := []int {-1, 0, 1, 2, -1, -4}
	r10 := [][]int {{-1, 0, 1}, {-1, -1, 2}}
	testCase(s1, r10, t)
}
