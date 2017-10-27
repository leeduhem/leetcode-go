package p077_1

func combine(n int, k int) [][]int {
	ns := make([]int, n)
	for i := 0; i < n; i++ {
		ns[i] = i + 1
	}

	combs := [][]int {}
	for i := 0; i < n-k+1; i++ {
		combs1 := [][]int { {ns[i]} }
		for j := 1; j < k; j++ {
			combs2 := [][]int {}
			for _, c := range combs1 {
				c1 := make([]int, len(c))
				copy(c1, c)

				combs3 := [][]int {}
				for l := c1[len(c1)-1]; l < n; l++ {
					combs3 = append(combs3, append(c1, ns[l]))
				}
				combs2 = append(combs2, combs3...)
			}
			combs1 = combs2
		}
		combs = append(combs, combs1...)
	}

	return combs
}
