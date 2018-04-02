package p207_2

func makeGraph(nodes int, edeges [][]int) [][]int {
	graph := make([][]int, nodes)
	for i, _ := range graph {
		graph[i] = make([]int, 0)
	}

	for _, edege := range edeges {
		graph[edege[1]] = append(graph[edege[1]], edege[0])
	}
	return graph
}

func computeIndegree(graph [][]int) []int {
	inDegrees := make([]int, len(graph))
	for _, neighbors := range graph {
		for _, neighbor := range neighbors {
			inDegrees[neighbor]++
		}
	}
	return inDegrees
}

func isCyclic(nodes int, edeges [][]int) bool {
	graph := makeGraph(nodes, edeges)
	inDegrees := computeIndegree(graph)
	for i := 0; i < nodes; i++ {
		j := 0
		for ; j < nodes; j++ {
			if inDegrees[j] == 0 {
				break
			}
		}
		if j == nodes {
			return true
		}

		inDegrees[j] = -1
		for _, neighbor := range graph[j] {
			inDegrees[neighbor]--
		}
	}

	return false
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	if isCyclic(numCourses, prerequisites) {
		return false
	}
	return true
}
