package p210_1

import (
	"errors"
)

func reverse(a []int) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		j := len(a) - 1 - i
		a[i], a[j] = a[j], a[i]
	}
}

type Graph struct {
	nodes         int
	adjacencyList [][]int
}

func makeGraph(nodes int, edges [][]int) *Graph {
	adjacencyList := make([][]int, nodes)
	for i := 0; i < nodes; i++ {
		adjacencyList[i] = make([]int, 0)
	}

	for _, edge := range edges {
		to, from := edge[0], edge[1]
		adjacencyList[from] = append(adjacencyList[from], to)
	}

	return &Graph{
		nodes:         nodes,
		adjacencyList: adjacencyList,
	}
}

// Topological sort by DFS
func topoSort(graph *Graph) ([]int, error) {
	visited := make(map[int]int)
	for i := 0; i < graph.nodes; i++ {
		// 0 -> unvisited, 1 -> visiting, 2 -> visited
		visited[i] = 0
	}

	rorder := make([]int, 0)

	var canSort func(from int) bool
	canSort = func(from int) bool {
		if visit := visited[from]; visit == 2 {
			return true
		} else if visit == 1 {
			return false
		}

		visited[from] = 1
		for _, to := range graph.adjacencyList[from] {
			if !canSort(to) {
				return false
			}
		}
		visited[from] = 2
		rorder = append(rorder, from)
		return true
	}

	for from := 0; from < graph.nodes; from++ {
		if !canSort(from) {
			return nil, errors.New("Failed to topological sort")
		}
	}

	reverse(rorder)
	return rorder, nil
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := makeGraph(numCourses, prerequisites)

	if order, err := topoSort(graph); err != nil {
		return []int{}
	} else {
		return order
	}
}
