package p088

func merge(nums1 []int, m int, nums2 []int, n int)  {
	nums := make([]int, m+n)

	i,j,k := 0,0,0
	for ; i < m && j < n; k++ {
		if nums1[i] < nums2[j] {
			nums[k] = nums1[i]; i++
		} else {
			nums[k] = nums2[j]; j++
		}
	}

	for ; i < m; i, k = i+1, k+1 {
		nums[k] = nums1[i]
	}

	for ; j < n; j, k = j+1, k+1 {
		nums[k] = nums2[j]
	}

	copy(nums1, nums)
}
