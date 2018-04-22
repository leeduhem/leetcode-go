// https://leetcode.com/problems/median-of-two-sorted-arrays/description/
package p004

func intMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func getKth(s, l []int, k int) int {
	m, n := len(s), len(l)

	// let m <= n
	if m > n {
		return getKth(l, s, k)
	}

	if m == 0 {
		return l[k-1]
	}
	if k == 1 {
		return intMin(s[0], l[0])
	}

	i := intMin(m, k/2)
	j := intMin(n, k/2)
	if s[i-1] > l[j-1] {
		return getKth(s, l[j:], k-j)
	}
	return getKth(s[i:], l, k-i)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	l := (m + n + 1) / 2
	r := (m + n + 2) / 2

	return (float64)(getKth(nums1, nums2, l)+getKth(nums1, nums2, r)) / 2
}
