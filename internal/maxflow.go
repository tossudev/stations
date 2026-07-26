package internal

// "we're Karping all over the place"

import (
	"slices"
	"strconv"
	"fmt"
)

// Edmonds-Karp
func MaxFlow(g *GraphList, start, end string) (int, [][]int) {

	// start from source(out)
	// NOTE: sink is assigned to sink(in) by default

	// matrix looks like this:
	//
	//		in		out
	// in	[in-in][in-out]
	// out	[out-in][out-out]
	//
	// naturally only in-out and out-in will be populated

	var matrixSize int = len(g.stationsNames)*2
	var source int = slices.Index(graphList.stationsNames, start) + matrixSize/2
	var sink int = slices.Index(graphList.stationsNames, end)

	// residual graph is a copy of our adjacency matrix
	residual := make([][]int, matrixSize)
	for i := range residual {
		residual[i] = make([]int, matrixSize)
		copy(residual[i], g.adjMatrix[i])
	}

	// parent saves bfs path source->sink
	parent := make([]int, matrixSize)
	maxFlow := 0

	// While there exists an augmenting path from source to sink
	for g.bfs(source, sink, parent, residual, matrixSize) {
		fmt.Println("Parent:", parent)
		// Update residual capacities
		for v := sink; v != source; v = parent[v] {
			u := parent[v]
			residual[u][v] -= 1 // Decrease forward edge
			residual[v][u] += 1 // Increase backward edge
		}
		maxFlow += 1
	}

	paths := findPathsFromResidual(residual, g.adjMatrix, source, sink, matrixSize/2)

	Log("Max flow:", strconv.Itoa(maxFlow))
	return maxFlow, paths
}


func findPathsFromResidual(residual [][]int, adjMatrix [][]int, source, sink, n int) [][]int {
	var paths [][]int

	for {
		current := source
		path := []int{source}
		path_found := false

		for current != sink {
			next_node := -1
			for v := range n * 2 {
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


// bfs performs breadth-first search to find a path from source to sink
// Returns true if a path exists, and fills the parent array with the path
func (g *GraphList) bfs(source, sink int, parent []int, residual [][]int, matrixSize int) bool {
	visited := make([]bool, matrixSize)

	// Create queue and add source
	queue := []int{source}
	visited[source] = true
	parent[source] = -1

	// Standard BFS
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		// Check all adjacent vertices
		// TODO: this is probably suboptimal but we let it sliiiiiide for now
		for v := 0; v < matrixSize; v++ {
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
