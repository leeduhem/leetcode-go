package p210

import (
	"errors"
)

type Graph struct {
	nodes int
	inDegrees []int
	adjacencyList [][]int
}

func makeGraph(nodes int, edges [][]int) *Graph {
	inDegrees := make([]int, nodes)
	adjacencyList := make([][]int, nodes)
	for i := 0; i < nodes; i++ {
		adjacencyList[i] = make([]int, 0)
	}

	for _, edge := range edges {
		to, from := edge[0], edge[1]
		inDegrees[to]++
		adjacencyList[from] = append(adjacencyList[from], to)
	}

	return &Graph{
		nodes : nodes,
		inDegrees: inDegrees,
		adjacencyList: adjacencyList,
	}
}

// Topological sort by BFS
func topoSort(graph *Graph) ([]int, error) {
	order := make([]int, graph.nodes)
	toVisit := make([]int, 0)

	for i, d := range graph.inDegrees {
		if d == 0 {
			toVisit = append(toVisit, i)
		}
	}

	visited := 0
	for len(toVisit) != 0 {
		from := toVisit[0]
		toVisit = toVisit[1:]

		order[visited] = from
		visited++

		for _, to := range graph.adjacencyList[from] {
			graph.inDegrees[to]--
			if graph.inDegrees[to] == 0 {
				toVisit = append(toVisit, to)
			}
		}
	}

	if visited == graph.nodes {
		return order, nil
	}
	return nil, errors.New("Failed to topological sort")
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := makeGraph(numCourses, prerequisites)

	if order, err := topoSort(graph); err != nil {
		return []int{}
	} else {
		return order
	}
}
