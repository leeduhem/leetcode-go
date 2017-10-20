package p049

import "sort"

func str2UniqueKey(s string) string {
	s1 := []rune(s)
	sort.Slice(s1, func (i, j int) bool { return s1[i] < s1[j] })
	return string(s1)
}

func groupAnagrams(strs []string) [][]string {
	m := map[string]([]string) {}

	for _,s := range strs {
		k := str2UniqueKey(s)
		if v,ok := m[k]; ok {
			m[k] = append(v, s)
		} else {
			m[k] = []string { s }
		}
	}

	res := [][]string {}
	for _,v := range m {
		res = append(res, v)
	}
	return res
}
