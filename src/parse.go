package main

import (
	"strings"
	"errors"
	"strconv"
	"fmt"
)

const (
	maxStations int = 10_000
)


func ParseMap(input string) ([]Station, []Connection, bool) {
	var stations []Station
	var connections []Connection

	var parseStations bool = false
	var parseConnections bool = false
	
	if !strings.Contains(input, "stations:") {
		PrintErr(ErrNoStations)
		return []Station{}, []Connection{}, false
	}
	
	if !strings.Contains(input, "connections:") {
		PrintErr(ErrNoConnections)
		return []Station{}, []Connection{}, false
	}

	for _, line := range strings.Split(input, "\n") {
		var commentIndex int = strings.Index(line, "#")
		if commentIndex != -1 {
			line = line[:commentIndex]
		}

		line = strings.ReplaceAll(line, " ", "")
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
			station, err := parseStation(line)
			if err != nil {
				return stations, connections, false
			}
			stations = append(stations, station)
		}
		if parseConnections {
			connection, err := parseConnection(line)
			if err != nil {
				return stations, connections, false
			}
			connections = append(connections, connection)
		}
	}

	if len(stations) > maxStations {
		PrintErr(ErrTooManyStations)
		return []Station{}, []Connection{}, false
	}

	return stations, connections, true
}


func parseStation(input string) (Station, error) {
	var stationValues []string = strings.Split(input, ",")
	if len(stationValues) != 3 {
		PrintErrArgs(ErrMalformedStation, input)
		return Station{}, errors.New(fmt.Sprintf(ErrMalformedStation, input))
	}
	
	name := stationValues[0]
	x, err := strconv.Atoi(stationValues[1])
	if err != nil || x < 0 {
		PrintErr(ErrInvalidCoordinates)
		return Station{}, errors.New(ErrInvalidCoordinates)
	}
	y, err := strconv.Atoi(stationValues[2])
	if err != nil || y < 0 {
		PrintErr(ErrInvalidCoordinates)
		return Station{}, errors.New(ErrInvalidCoordinates)
	}

	return Station{
		Name:	name,
		X:		x,
		Y:		y,
	}, nil
}


func parseConnection(input string) (Connection, error) {
	var connectionValues []string = strings.Split(input, "-")
	if len(connectionValues) != 2 {
		PrintErr(ErrMalformedConnection, input)
		return Connection{}, errors.New(fmt.Sprintf(ErrMalformedConnection, input))
	}

	return Connection{
		Begin:	connectionValues[0],
		End:	connectionValues[1],
	}, nil
}
