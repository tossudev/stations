package main

import (
	"math"
	"fmt"
)


// MaxFlow calculates the maximum flow from source to sink using Edmonds-Karp
func MaxFlow(g *GraphList, source, sink int) int {
    // Create residual graph - initially same as capacity
    residual := make([][]int, len(g.stationsNames))
    for i := range residual {
        residual[i] = make([]int, len(g.stationsNames))
        copy(residual[i], g.adjMatrix[i])
    }
    
    parent := make([]int, len(g.stationsNames))
    maxFlow := 0
    
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
        
        // Update residual capacities
        for v := sink; v != source; v = parent[v] {
            u := parent[v]
            residual[u][v] -= pathFlow  // Decrease forward edge
            residual[v][u] += pathFlow  // Increase backward edge
        }
        maxFlow += pathFlow
		fmt.Println("NEW")
		fmt.Println(parent)
		for child, value := range parent {
			if value == 0 {
				continue
			}

			if value == -1 {
				fmt.Println("connection:", g.stationsNames[child + 2], "->", g.stationsNames[0])
			continue
			}

			fmt.Println("connection:", g.stationsNames[child], "->", g.stationsNames[value])
		}
    }


    return maxFlow
}


// bfs performs breadth-first search to find a path from source to sink
// Returns true if a path exists, and fills the parent array with the path
func (g *GraphList) bfs(source, sink int, parent []int, residual [][]int) bool {
    // Initialize visited array
    visited := make([]bool, len(g.stationsNames))
    
    // Create queue and add source
    queue := []int{source}
    visited[source] = true
    parent[source] = -1
    
    // Standard BFS
    for len(queue) > 0 {
        u := queue[0]
        queue = queue[1:]
        
        // Check all adjacent vertices
        for v := 0; v < len(g.stationsNames); v++ {
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

