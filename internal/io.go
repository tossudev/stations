package internal

import (
	"os"
	"strconv"
	"slices"
)

const (
	maxTrains	int = 100_000
)

func ParseArgs() (mapfile, start, end string, trainCount int, ok bool) {
	ok = false

	if slices.Contains(os.Args, "-h") || slices.Contains(os.Args, "--help") {
		PrintUsage()
		os.Exit(0)
	}

	if len(os.Args) < 5 {
		PrintErr(ErrArgsCount)
		PrintUsage()
		return
	}

	// find verbose flag and remove it to work with indices in order
	// maybe using the flag package would be cleaner, this works fine
	if len(os.Args) == 6 {
		if slices.Contains(os.Args, "-v") || slices.Contains(os.Args, "--verbose") {
			Verbose = true
			
			i := slices.Index(os.Args, "-v")
			if i == -1 {
				i = slices.Index(os.Args, "--verbose")
			}
			
			os.Args = append(os.Args[:i], os.Args[i+1:]...)
		}
	}

	mapfile = os.Args[1]
	start = os.Args[2]
	end = os.Args[3]
	trainCountStr := os.Args[4]
	
	Log("Using", mapfile, "as input.")

	var err error
	trainCount, err = strconv.Atoi(trainCountStr)
	if err != nil || trainCount < 1 {
		PrintErr(ErrTrainsCount)
		return
	}

	if trainCount > maxTrains {
		PrintErr(ErrTooManyTrains)
		return
	}

	if start == end {
		PrintErr(ErrStationsSame)
		return
	}

	ok = true
	return
}

func ReadMapFile(filename string) string {
	contents, err := os.ReadFile(filename)
	if err != nil {
		PrintErr(err.Error())
		return ""
	}
	return string(contents)
}
