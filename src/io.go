package main

import (
	"os"
)


func ParseArgs() (mapfile, outputfile string, ok bool) {
	ok = false

	if len(os.Args) != 3 {
		PrintErr(ErrArgsCount)
		Log(Usage)
		return
	}

	mapfile = os.Args[1]
	outputfile = os.Args[2]
	
	Log("Using", mapfile, "as input.")

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
