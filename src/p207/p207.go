package p207

func isCyclic(edeges [][]int) bool {
	em := map[int]int{}
	for _, e := range edeges {
		em[e[0]] = e[1]
	}

	for _, e := range edeges {
		f, s := e[0], e[1]
		node := map[int]bool{}
		for {
			if n, ok := em[s]; ok {
				if _, ok := node[n]; ok {
					return true
				} else {
					node[n] = true
				}

				if n == f {
					return true
				}
				s = n
			} else {
				break
			}
		}
	}

	return false
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	if isCyclic(prerequisites) {
		return false
	}
	return true
}
