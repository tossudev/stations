package internal

import (
	"fmt"
	"slices"
)

type Station struct {
	Name string
	X    int
	Y    int
}

type GraphList struct {
	stations      map[string]*Station
	coordinateMap map[string]string
	stationsNames []string

	// adjacency matrix represents connections
	// it looks like this:
	//
	//		in		out
	// in	[in-in][in-out]
	// out	[out-in][out-out]
	//
	// naturally only in-out and out-in will be populated
	adjMatrix     [][]int
}


func NewGraphlist() *GraphList {
	return &GraphList{
		stations:      make(map[string]*Station),
		coordinateMap: make(map[string]string),
		stationsNames: []string{},
	}
}

func (g *GraphList) AddStation(station *Station) bool {
	if _, exists := g.stations[station.Name]; exists {
		PrintErrArgs(ErrDuplicateStations, station.Name)
		return false
	}
	
	coords := fmt.Sprintf("%d,%d", station.X, station.Y)
	if existingName, exists := g.coordinateMap[coords]; exists {
		PrintErrArgs(ErrDuplicateCoordinates, coords, "stations: ", existingName, station.Name)
		return false
	}

	g.stations[station.Name] = station
	g.coordinateMap[coords] = station.Name
	g.stationsNames = append(g.stationsNames, station.Name)

	return true
}

func (g *GraphList) AddConnection(from, to string) bool {
	stationsCount := len(g.stationsNames)
	
	// create adjacency matrix if it does not exist
	if len(g.adjMatrix) == 0 {
		g.adjMatrix = make([][]int, stationsCount*2)
		for i := range g.adjMatrix {
			g.adjMatrix[i] = make([]int, stationsCount*2)
		}
	}

	if _, exists := g.stations[from]; !exists {
		PrintErrArgs(ErrStationNotExist, from)
		return false
	}
	if _, exists := g.stations[to]; !exists {
		PrintErrArgs(ErrStationNotExist, to)
		return false
	}

	nfrom := slices.Index(g.stationsNames, from)
	nto := slices.Index(g.stationsNames, to)

	nfrom_in := nfrom
	nfrom_out := nfrom + stationsCount

	nto_in := nto
	nto_out := nto + stationsCount

	if g.adjMatrix[nfrom_out][nto_in] == 1 {
		PrintErrArgs(ErrDuplicateConnections, to, from)
		return false
	}

	// add bidirected connection between stations, eg. waterloo(out) -> euston(in)
	g.adjMatrix[nfrom_out][nto_in] = 1
	g.adjMatrix[nto_out][nfrom_in] = 1

	// add split node bidirected connections, eg. waterloo(in) -> waterloo(out)
	g.adjMatrix[nfrom_in][nfrom_out] = 1
	g.adjMatrix[nfrom_out][nfrom_in] = 1

	g.adjMatrix[nto_in][nto_out] = 1
	g.adjMatrix[nto_out][nto_in] = 1

	return true
}
