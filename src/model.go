package main

import (
	"fmt"
	"slices"
)

type Station struct {
	Name string
	X    int
	Y    int
}
type Graph struct {
	vertices int
	capacity [][]int
}

func NewGraph(vertices int) *Graph {
	capacity := make([][]int, vertices)
	for i := range capacity {
		capacity[i] = make([]int, vertices)
	}
	return &Graph{
		vertices: vertices,
		capacity: capacity,
	}
}

func (g *Graph) AddEdge(u, v, cap int) {
	g.capacity[u][v] = cap
}

type GraphList struct {
	stations      map[string]*Station
	adjacentList  map[string][]string
	coordinateMap map[string]string
}

func NewGraphlist() *GraphList {
	return &GraphList{
		stations:      make(map[string]*Station),
		adjacentList:  make(map[string][]string),
		coordinateMap: make(map[string]string),
	}
}

func (g *GraphList) AddStation(station *Station) bool {
	if _, exists := g.stations[station.Name]; exists {
		PrintErrArgs(ErrDuplicateStations, station.Name)
		return false
	}
	// check dublicate coordinates
	coordKey := fmt.Sprintf("%d,%d", station.X, station.Y)
	if existingName, exists := g.coordinateMap[coordKey]; exists {
		PrintErrArgs(ErrDuplicateCoordinates, coordKey, "stations: ", existingName, station.Name)
		return false
	}
	// all checks pass
	g.stations[station.Name] = station
	g.adjacentList[station.Name] = []string{}
	g.coordinateMap[coordKey] = station.Name
	return true
}

// dont know if these errors ar correct?
func (g *GraphList) AddConnection(from, to string) bool {
	if _, exists := g.stations[from]; !exists {
		PrintErrArgs(ErrStationNotExist, from)
		return false
	}
	if _, exists := g.stations[to]; !exists {
		PrintErrArgs(ErrStationNotExist, to)
		return false
	}
	if slices.Contains(g.adjacentList[from], to) {
		PrintErrArgs(ErrDuplicateConnections, to, from)
		return false
	}

	g.adjacentList[from] = append(g.adjacentList[from], to)
	g.adjacentList[to] = append(g.adjacentList[to], from)
	return true
}
