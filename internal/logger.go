package internal

import (
	"fmt"
	"os"
	"strings"
)

const (
	PrefixError   string = "\u001B[31m[ERROR]\t"
	PrefixWarning string = "\u001B[33m[WARN]\t"
	SuffixReset   string = "\033[0m"
	PrefixLog     string = "\u001B[32m[LOG]\t"

	Usage                    string = `
Usage: ./stations [MAP FILE] [START] [END] [TRAINS AMOUNT] [optional args...]

Optional arguments:
	-v, --verbose
		Verbose output

	-h, --help
		Show this help message`
	ErrArgsCount             string = "Incorrect number of command line arguments"
	ErrStartStationNotExist  string = "Start station does not exist"
	ErrEndStationNotExist    string = "End station does not exist"
	ErrStationsSame          string = "Start and end station are the same"
	ErrNoPath                string = "No path between the start and end stations"
	ErrDuplicateConnections  string = "Duplicate connections: %s"
	ErrTrainsCount           string = "Number of trains is not a valid positive integer"
	ErrInvalidCoordinates    string = "Coordinates are not valid positive integers"
	ErrDuplicateCoordinates  string = "Two stations exist at the exact same coordinate location: %s"
	ErrConnectionNotExist    string = "Connection does not exist: %s"
	ErrDuplicateStations     string = "Duplicate station names: %s"
	ErrInvalidStationNameArg string = "Invalid station name in argument"
	ErrNoStations            string = "Map does not contain stations section"
	ErrNoConnections         string = "Map does not contain connections section"
	ErrTooManyStations       string = "Station limit is 10,000"
	ErrTooManyTrains       	 string = "Train limit is 100,000"
	ErrMalformedStation      string = "Malformed station data: %s"
	ErrMalformedConnection   string = "Malformed connection data: %s"
	ErrStationNotExist       string = "Station does not exist: %s"
)

var (
	Verbose	bool = false
)

func PrintErr(messages ...string) {
	message := strings.Join(messages, " ")
	message = fmt.Sprintf("%s%s%s\n", PrefixError, message, SuffixReset)
	os.Stderr.Write([]byte(message))
}

// Prints one of the defined messages with string format

// Example:
// PrintErrArgs(ErrDuplicateStation, "hakaniemi", "kalasatama")
// Will print out:
// [ERROR] Duplicate station names: [hakaniemi kalasatama]
func PrintErrArgs(errMessage string, args ...string) {
	message := fmt.Sprintf("%s%s%s\n", PrefixError, errMessage, SuffixReset)
	message = fmt.Sprintf(message, args)
	os.Stderr.Write([]byte(message))
}

func Log(messages ...string) {
	if Verbose {
		message := strings.Join(messages, " ")
		fmt.Printf("%s%s%s\n", PrefixLog, message, SuffixReset)
	}
}

func PrintUsage() {
	fmt.Println(Usage)
	fmt.Println()
}

func PrintWarn(messages ...string) {
	message := strings.Join(messages, " ")
	fmt.Printf("%s%s%s\n", PrefixWarning, message, SuffixReset)
}
