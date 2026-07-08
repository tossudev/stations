package main

import (
	"fmt"
	"os"
)

var (
	ErrArgsCount 				string = "Incorrect number of command line arguments"
	ErrStartStationNotExist 	string = "Start station does not exist"
	ErrEndStationNotExist		string = "End station does not exist"
	ErrStationsSame				string = "Start and end station are the same"
	ErrNoPath					string = "No path between the start and end stations"
	ErrDuplicateConnections		string = "Duplicate connections: %s"
	ErrTrainsCount				string = "Number of trains is not a valid positive integer"
	ErrInvalidCoordinates		string = "Coordinates are not valid positive integers"
	ErrDuplicateCoordinates		string = "Two stations exist at the exact same coordinate location: %s"
	ErrConnectionNotExist		string = "Connection does not exist: %s"
	ErrDuplicateStation			string = "Duplicate station names: %s"
	ErrInvalidStationNameArg	string = "Invalid station name in argument"
	ErrNoStations				string = "Map does not contain stations section"
	ErrNoConnections			string = "Map does not contain connections section"
	ErrTooManyStations			string = "Map has over 10 000 stations"
)

func main() {
	PrintErrArgs(ErrDuplicateConnections, []string{"station1", "station2"})
	PrintErr(ErrArgsCount)
}

func PrintErr(name string) {
	err := "ERROR: " + name + "\n"
	os.Stderr.Write([]byte(err))
}

func PrintErrArgs(name string, args []string) {
	err := "ERROR: " + fmt.Sprintf(name, args) + "\n"
	os.Stderr.Write([]byte(err))
}

