package main

import "fmt"

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
		Log(ErrDuplicateStation, station.Name)
		PrintErrArgs(ErrDuplicateStation, station.Name)
		return false
	}
	// check dublicate coordinates
	coordKey := fmt.Sprintf("%d,%d", station.X, station.Y)
	if existingName, exists := g.coordinateMap[coordKey]; exists {
		Log(ErrDuplicateCoordinates, coordKey, "stations: ", existingName, station.Name)
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
		Log(ErrStartStationNotExist, from)
		PrintErrArgs(ErrStartStationNotExist, from)
		return false
	}
	if _, exists := g.stations[to]; !exists {
		Log(ErrEndStationNotExist, to)
		PrintErrArgs(ErrEndStationNotExist, to)
		return false
	}
	for _, neighbor := range g.adjacentList[from] {
		if neighbor == to {
			Log(ErrDuplicateConnections, from, to)
			PrintErrArgs(ErrDuplicateConnections, to, from)
			return false
		}
	}

	g.adjacentList[from] = append(g.adjacentList[from], to)
	g.adjacentList[to] = append(g.adjacentList[to], from)
	return true
}
