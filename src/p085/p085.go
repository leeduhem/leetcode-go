package p085

func andOfBins(b1, b2 []byte) []byte {
	if len(b1) != len(b2) {
		panic("Invalid input")
	}

	b := make([]byte, len(b1))
	for i := 0; i < len(b1); i++ {
		if b1[i] == '1' && b2[i] == '1' {
			b[i] = '1'
		} else {
			b[i] = '0'
		}
	}

	return b
}

func successiveOnes(b []byte) int {
	ones, cnt := 0, 0
	for _, c := range b {
		if c == '0' {
			if ones < cnt {
				ones = cnt
			}
			cnt = 0
			continue
		}
		cnt++
	}
	if ones < cnt {
		ones = cnt
	}
	return ones
}

func maximalRectangle(matrix [][]byte) int {
	maxArea := 0
	for i := 0; i < len(matrix); i++ {
		acc := matrix[i]
		for h, j := 1, i; j < len(matrix); h, j = h+1, j+1 {
			acc = andOfBins(acc, matrix[j])
			area := h * successiveOnes(acc)
			if maxArea < area {
				maxArea = area
			}
		}
	}

	return maxArea
}
