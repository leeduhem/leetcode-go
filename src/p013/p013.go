package p013

func baseRoman2Int(n byte) int {
	switch n {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}

	return 0
}

func romanToInt(s string) int {
	sl := len(s)
	num := 0

	for i := 0; i < sl; {
		n1 := baseRoman2Int(s[i])
		if i + 1 < sl {
			n2 := baseRoman2Int(s[i+1])
			if n1 < n2 {
				num += n2 - n1
				i += 2
				continue
			}

		}

		num += n1
		i += 1
	}

	return num
}
