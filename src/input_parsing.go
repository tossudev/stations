package main

import (
	"fmt"
	"strings"
	"errors"
	"strconv"
)


func ParseInput(input string) ([]Station, []Connection, error) {
	var parseStations bool = false
	var parseConnections bool = false

	var stations []Station
	var connections []Connection
	
	for i, line := range strings.Split(input, "\n") {
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
				fmt.Println("ERROR parsing line", i, "\n", line)
				return stations, connections, err
			}
			stations = append(stations, station)
		}
		if parseConnections {
			connection, err := parseConnection(line)
			if err != nil {
				fmt.Println("ERROR parsing line", i, "\n", line)
				return stations, connections, err
			}
			connections = append(connections, connection)
		}
	}
	
	return stations, connections, nil
}


func parseStation(input string) (Station, error) {
	var stationValues []string = strings.Split(input, ",")
	if len(stationValues) != 3 {
		return Station{}, errors.New("Input does not have 3 values for station!")
	}
	
	name := stationValues[0]
	x, err := strconv.Atoi(stationValues[1])
	if err != nil {
		return Station{}, errors.New("Failed converting x to int!")
	}
	y, err := strconv.Atoi(stationValues[2])
	if err != nil {
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
		return Connection{}, errors.New("Input does not have 2 values for connection!")
	}

	return Connection{
		Begin:	connectionValues[0],
		End:	connectionValues[1],
	}, nil
}
