package main

import (
	"os"

	"stations/internal"
)


func main() {
	mapfile, start, end, trainCount, ok := internal.ParseArgs()
	if !ok {
		ExitWithErrors()
	}

	contents := internal.ReadMapFile(mapfile)

	graphList, ok := internal.ParseMap(string(contents), start, end)
	if !ok {
		ExitWithErrors()
	}

	_, paths := internal.MaxFlow(graphList, start, end)
	internal.CreateSchedule(graphList, paths, start, end, trainCount, true)
}


func ExitWithErrors() {
	internal.PrintWarn("Program exited with errors")
	os.Exit(0)
}
