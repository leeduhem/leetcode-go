// https://leetcode.com/problems/zigzag-conversion/description/
package p006

func convert(s string, numRows int) string {
	m := make([][]byte, numRows)

	len0 := len(s)
	isShort := false
	for i := 0; i < len0; {
		if !isShort {
			isShort = true

			for j := 0; j < numRows && i < len0; j++ {
				m[j] = append(m[j], s[i])
				i++
			}
		} else {
			isShort = false

			m[0] = append(m[0], 0)

			for j := numRows - 2; j > 0 && i < len0; j-- {
				m[j] = append(m[j], s[i])
				i++
			}

			if i < len0 {
				m[numRows-1] = append(m[numRows-1], 0)
			}
		}
	}

	result := make([]byte, 0)
	for _, r := range m {
		for _, c := range r {
			if c != 0 {
				result = append(result, c)
			}
		}
	}

	return string(result)
}
