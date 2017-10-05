package p017

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string {}
	}

	d2s := map[int]string {
		0 : "",
		1 : "",
		2 : "abc",
		3 : "def",
		4 : "ghi",
		5 : "jkl",
		6 : "mno",
		7 : "pqrs",
		8 : "tuv",
		9 : "wxyz",
	}

	ss := []string {""}
	for _,d := range digits {
		d1 := d - '0'
		ds := d2s[int(d1)]
		if ds == "" {
			continue
		}

		dslen := len(ds)
		ss1 := make([]string, len(ss) * dslen)
		for i,s := range ss {
			for j,c := range ds {
				ss1[i * dslen + j] = s + string([]rune{c})
			}
		}

		ss = ss1
	}

	return ss
}
