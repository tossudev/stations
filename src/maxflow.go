package main

import (
	"math"
	"fmt"
)

// MaxFlow calculates the maximum flow from source to sink using Edmonds-Karp
func MaxFlow(g *GraphList, source, sink int) (int, [][]int) {
	// Create residual graph - initially same as capacity
	residual := make([][]int, len(g.stationsNames)*2)
	for i := range residual {
		residual[i] = make([]int, len(g.stationsNames)*2)
		copy(residual[i], g.adjMatrix[i])
	}

	for _, val := range residual {
		fmt.Println(val)
	}

	parent := make([]int, len(g.stationsNames)*2)
	maxFlow := 0
	var paths [][]int
	source += len(g.stationsNames)

	// While there exists an augmenting path from source to sink
	for g.bfs(source, sink, parent, residual) {

		// Find minimum residual capacity along the path
		pathFlow := math.MaxInt32
		for v := sink; v != source; v = parent[v] {
			u := parent[v]
			if residual[u][v] < pathFlow {
				pathFlow = residual[u][v]
			}
		}

		path := []int{}
		for v := sink; v != source; v = parent[v] {
			path = append(path, v)
			// v = parent[v] // move to next node in path
		}


		path = append(path, source)
		reverse(path)
		paths = append(paths, path)

		// Update residual capacities
		for v := sink; v != source; v = parent[v] {
			u := parent[v]
			residual[u][v] -= pathFlow // Decrease forward edge
			residual[v][u] += pathFlow // Increase backward edge
		}
		maxFlow += pathFlow

		/*
		fmt.Println("NEW")
		fmt.Println(parent)
		for child, value := range parent {
			if value == 0 {
				continue
			}

			child = child % len(g.stationsNames)
			value = value % len(g.stationsNames)

			if value == -1 {
				fmt.Println("connection:", g.stationsNames[child+2], "->", g.stationsNames[0])
				continue
			}

			fmt.Println("connection:", g.stationsNames[child], "->", g.stationsNames[value])
		}
		*/
	}

	newPaths := findPathsFromResidual(residual, g.adjMatrix, source, sink, len(g.stationsNames))

	return maxFlow, newPaths
}


func findPathsFromResidual(residual [][]int, adjMatrix [][]int, source, sink, n int) [][]int {
	var paths [][]int

	for {
		current := source
		path := []int{source}
		path_found := false

		for current != sink {
			next_node := -1
			for v := range n*2 {
				if current < n && v == current + n && residual[current][v] == 0 || current >= n && v < n && adjMatrix[current][v] == 1 && residual[current][v] == 0 {
					next_node = v
					break
				}
			}

			if next_node == -1 {
				break
			}
			
			residual[current][next_node] = 1

			current = next_node
			if current < n {
				path = append(path, current)
			}
			if current == sink {
				path_found = true
				break
			}
		}

		if path_found {
			paths = append(paths, path)
		} else {
			break
		}
	}

	return paths
}


// helper to reverse path
func reverse(path []int) {
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
}

// bfs performs breadth-first search to find a path from source to sink
// Returns true if a path exists, and fills the parent array with the path
func (g *GraphList) bfs(source, sink int, parent []int, residual [][]int) bool {
	// Initialize visited array
	visited := make([]bool, len(g.stationsNames)*2)

	// Create queue and add source
	queue := []int{source}
	visited[source] = true
	parent[source] = -1

	// Standard BFS
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		// Check all adjacent vertices
		for v := 0; v < len(g.stationsNames)*2; v++ {
			// If not visited and has residual capacity
			if !visited[v] && residual[u][v] > 0 {
				queue = append(queue, v)
				parent[v] = u
				visited[v] = true

				// Early exit if we reached the sink
				if v == sink {
					return true
				}
			}
		}
	}

	return false
}
