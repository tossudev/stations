package main

type Station struct {
	Name string
	X    int
	Y    int
}

// // TODO: Stations could be stored as actual Station struct instead of name string
//
//	type Connection struct {
//		Begin string
//		End   string
//	}
type GraphList struct {
	stations map[string]*Station
	adjList  map[string][]string
}

func NewGraphlist() *GraphList {
	return &GraphList{
		adjList: make(map[string][]string),
	}
}

// add stuff
func (g *GraphList) AddStation(station string) {
	if _, exists := g.adjList[station]; !exists {
		g.adjList[station] = []string{}
	}
}

func (g *GraphList) AddConnection(from, to string) {
	g.AddStation(from)
	g.AddStation(to)
	g.adjList[from] = append(g.adjList[from], to)
	g.adjList[to] = append(g.adjList[to], from)
}
