package p207_1

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

func isCyclic(nodes int, edeges [][]int) bool {
	graph := makeGraph(nodes, edeges)
	onPath := make([]bool, nodes)
	visited := make([]bool, nodes)

	var dfsIsCyclic func(node int) bool
	dfsIsCyclic = func(node int) bool {
		if visited[node] {
			return false
		}
		onPath[node], visited[node] = true, true
		for _, neighbor := range graph[node] {
			if onPath[neighbor] || dfsIsCyclic(neighbor) {
				return true
			}
		}
		onPath[node] = false

		return false
	}

	for i := 0; i < nodes; i++ {
		if !visited[i] && dfsIsCyclic(i) {
			return true
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
