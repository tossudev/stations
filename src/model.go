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
		stations: make(map[string]*Station),
		adjList:  make(map[string][]string),
	}
}

// add stuff
func (g *GraphList) AddStation(station *Station) {
	if _, exists := g.stations[station.Name]; !exists {
		g.stations[station.Name] = station
		g.adjList[station.Name] = []string{}
	}
}

func (g *GraphList) AddConnection(from, to string) {
	if _, exists := g.stations[from]; !exists {
		return // error
	}
	if _, exists := g.stations[to]; !exists {
		return // error
	}

	g.adjList[from] = append(g.adjList[from], to)
	g.adjList[to] = append(g.adjList[to], from)
}
