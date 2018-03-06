package p165

import (
	"strings"
	"strconv"
)

func intMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func compareVersion(version1 string, version2 string) int {
	vers1 := strings.Split(version1, ".")
	vers2 := strings.Split(version2, ".")
	l1, l2 := len(vers1), len(vers2)

	for i := 0; i < intMin(l1, l2); i++ {
		v1, _ := strconv.Atoi(vers1[i])
		v2, _ := strconv.Atoi(vers2[i])
		if v1 > v2 {
			return 1
		} else if v1 < v2 {
			return -1
		}
	}

	if l1 > l2 {
		for i := l2; i < l1; i++ {
			v1, _ := strconv.Atoi(vers1[i])
			if v1 != 0 {
				return 1
			}
		}
	}

	if l1 < l2 {
		for i := l1; i < l2; i++ {
			v2, _ := strconv.Atoi(vers2[i])
			if v2 != 0 {
				return -1
			}
		}
	}

	return 0
}
