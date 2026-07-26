package main

import "os"


func main() {
	mapfile, start, end, trainCount, ok := ParseArgs()
	if !ok {
		ExitWithErrors()
	}

	contents := ReadMapFile(mapfile)

	graphList, ok := ParseMap(string(contents), start, end)
	if !ok {
		ExitWithErrors()
	}

	_, paths := MaxFlow(graphList, start, end)
	CreateSchedule(graphList, paths, start, end, trainCount, true)
}


func ExitWithErrors() {
	PrintWarn("Program exited with errors")
	os.Exit(0)
}
