package main


import (
	"fmt"
)


func main() {
	mapfile, start, end, trainCount, ok := ParseArgs()
	if !ok {
		Log("Program exited with errors.")
		return
	}

	contents := ReadMapFile(mapfile)

	graphList, ok := ParseMap(string(contents))
	if !ok {
		Log("Program exited with errors.")
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

	/*
	for name, station := range graphList.stations {
		fmt.Println(name, station)
	}
	for name2, station2 := range graphList.adjacentList {
		fmt.Println(name2, station2)
	}
	*/
	fmt.Println(mapfile, start, end, trainCount)
	
	fmt.Println(Bfs(graphList, start, end))
}
