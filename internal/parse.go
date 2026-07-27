package internal

import (
	"strconv"
	"fmt"
	"strings"
)

const (
	maxStations	int = 10_000
)


func ParseMap(input, start, end string) (*GraphList, bool) {
	graphList := NewGraphlist()

	var parseStations bool = false
	var parseConnections bool = false

	if !strings.Contains(input, "stations:") {
		PrintErr(ErrNoStations)
		return graphList, false
	}

	if !strings.Contains(input, "connections:") {
		PrintErr(ErrNoConnections)
		return graphList, false
	}

	for line := range strings.SplitSeq(input, "\n") {
		var commentIndex int = strings.Index(line, "#")
		if commentIndex != -1 {
			line = line[:commentIndex]
		}

		line = strings.ReplaceAll(line, " ", "")
		line = strings.ReplaceAll(line, "\r", "")
		line = strings.ReplaceAll(line, "\t", "")
		line = strings.ReplaceAll(line, "\n", "")
		line = strings.ReplaceAll(line, "\v", "")

		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "stations:") {
			parseStations = true
			continue
		}

		if strings.HasPrefix(line, "connections:") {
			parseStations = false
			parseConnections = true
			continue
		}

		if parseStations {
			station, ok := parseStation(line)
			if !ok {
				return graphList, false
			}
			ok = graphList.AddStation(&station)
			if !ok {
				return graphList, false
			}
		}
		if parseConnections {
			begin, end, ok := parseConnection(line)
			if !ok {
				return graphList, false
			}
			ok = graphList.AddConnection(begin, end)
			if !ok {
				return graphList, false
			}
		}
	}

	if len(graphList.stations) > maxStations {
		PrintErr(ErrTooManyStations)
		return graphList, false
	}

	if _, exists := graphList.stations[start]; !exists {
		PrintErr(ErrStartStationNotExist)
		return graphList, false
	}
	if _, exists := graphList.stations[end]; !exists {
		PrintErr(ErrEndStationNotExist)
		return graphList, false
	}

	return graphList, true
}

func parseStation(input string) (Station, bool) {
	var stationValues []string = strings.Split(input, ",")
	if len(stationValues) != 3 {
		PrintErrArgs(ErrMalformedStation, input)
		return Station{}, false
	}

	name := stationValues[0]
	x, err := strconv.Atoi(stationValues[1])
	if err != nil || x < 0 {
		fmt.Println(err)
		PrintErrArgs(ErrInvalidCoordinates, name)
		return Station{}, false
	}
	y, err := strconv.Atoi(stationValues[2])
	if err != nil || y < 0 {
		fmt.Println(err)
		PrintErrArgs(ErrInvalidCoordinates, name)
		return Station{}, false 
	}

	return Station{
		Name: name,
		X:    x,
		Y:    y,
	}, true
}

func parseConnection(input string) (string, string, bool) {
	var connectionValues []string = strings.Split(input, "-")
	if len(connectionValues) != 2 {
		PrintErr(ErrMalformedConnection, input)
		return "", "", false
	}

	return connectionValues[0], connectionValues[1], true
}
