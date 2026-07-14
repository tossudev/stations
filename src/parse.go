package main

import (
	"strings"
	"errors"
	"strconv"
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
		return stations, connections, false
	}
	
	if !strings.Contains(input, "connections:") {
		PrintErr(ErrNoConnections)
		return stations, connections, false
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

	return stations, connections, true
}

func parseStation(input string) (Station, error) {
	var stationValues []string = strings.Split(input, ",")
	if len(stationValues) != 3 {
		PrintErrArgs(ErrMalformedStation, input)
		return Station{}, errors.New("Input does not have 3 values for station!")
	}
	
	name := stationValues[0]
	x, err := strconv.Atoi(stationValues[1])
	if err != nil || x < 0{
		PrintErr(ErrInvalidCoordinates)
		return Station{}, errors.New("Failed converting x to int!")
	}
	y, err := strconv.Atoi(stationValues[2])
	if err != nil || y < 0{
		PrintErr(ErrInvalidCoordinates)
		return Station{}, errors.New("Failed converting y to int!")
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
		return Connection{}, errors.New("Input does not have 2 values for connection!")
	}

	return Connection{
		Begin:	connectionValues[0],
		End:	connectionValues[1],
	}, nil
}
