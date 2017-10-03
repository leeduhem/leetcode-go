package p004

func min(a, b int) int {
	if (a < b) {
		return a
	} else {
		return b
	}
}

func getkth(s, l []int, k int) int {
	m,n := len(s), len(l)

	// let m <= n
	if (m > n) {
		return getkth(l, s, k)
	}

	if (m == 0) {
		return l[k-1]
	}
	if (k == 1) {
		return min(s[0], l[0])
	}

	i := min(m, k/2)
	j := min(n, k/2)
	if (s[i-1] > l[j-1]) {
		return getkth(s, l[j:], k-j)
	} else {
		return getkth(s[i:], l, k-i)
	}
	return 0
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m,n := len(nums1), len(nums2)
	l := (m + n + 1) / 2
	r := (m + n + 2) / 2

	return (float64) (getkth(nums1, nums2, l) + getkth(nums1, nums2, r)) / 2
}
