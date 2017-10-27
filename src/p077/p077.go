package p077

func genCombine(ns []int, n, k int) [][]int {
	if k == 0 || k > n {
		return [][]int {{}}
	}
	if k == n {
		return [][]int {ns}
	}

	combs := [][]int {}
	for i := 0; i < len(ns) && i < n-k+1; i++ {
		combs1 := genCombine(ns[i+1:], n-1, k-1)
		for j, c := range combs1 {
			combs1[j] = append(c, ns[i])
		}
		combs = append(combs, combs1...)
	}

	return combs
}

func combine(n int, k int) [][]int {
	ns := make([]int, n)
	for i := 0; i < n; i++ {
		ns[i] = i + 1
	}

	return genCombine(ns, n, k)
}
