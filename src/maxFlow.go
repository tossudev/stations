package main

import "math"

func (g *Graph) Bfs(source, sink int, parent []int, residual [][]int) bool {
	visited := make([]bool, g.vertices)

	queue := []int{source}
	visited[source] = true
	parent[source] = -1

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		for v := 0; v < g.vertices; v++ {
			// if not visited and has residual capacity
			if !visited[v] && residual[u][v] > 0 {
				queue = append(queue, v)
				parent[v] = u
				visited[v] = true
				if v == sink {
					return true
				}
			}
		}
	}
	return false
}

func (g *Graph) MaxFlow(source, sink int) int {
	// residual graph, kinda same as capacity
	residual := make([][]int, g.vertices)
	for i := range residual {
		residual[i] = make([]int, g.vertices)
		copy(residual[i], g.capacity[i])
	}

	parent := make([]int, g.vertices)
	maxFlow := 0

	for g.Bfs(source, sink, parent, residual) {
		// find minimum residual capacity along the path
		pathFlow := math.MaxInt32
		for v := sink; v != source; v = parent[v] {
			u := parent[v]
			if residual[u][v] < pathFlow {
				pathFlow = residual[u][v]
			}
		}
		// update residual capacity
		for v := sink; v != source; v = parent[v] {
			u := parent[v]
			residual[u][v] -= pathFlow
			residual[v][u] += pathFlow
		}
		maxFlow += pathFlow
	}
	return maxFlow
}
