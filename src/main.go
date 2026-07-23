package main

import (
	"fmt"
)

func main() {
	mapfile, start, end, trainCount, ok := ParseArgs()
	if !ok {
		PrintWarn("Program exited with errors.")
		return
	}

	contents := ReadMapFile(mapfile)

	graphList, ok := ParseMap(string(contents))
	if !ok {
		PrintWarn("Program exited with errors.")
		return
	}

	if _, exists := graphList.stations[start]; !exists {
		PrintErr(ErrStartStationNotExist)
		return
	}
	if _, exists := graphList.stations[end]; !exists {
		PrintErr(ErrEndStationNotExist)
		return
	}
	
	mf, paths := MaxFlow(graphList, start, end)
	fmt.Println("Max flow:", mf)
	fmt.Println("Trains:", trainCount)

	for i, path := range paths {
		fmt.Printf("Path #%d: ", i + 1)
		for _, station := range path {
			fmt.Printf("%s ", graphList.stationsNames[station % len(graphList.stationsNames)])
		}
		fmt.Println()
	}
}
