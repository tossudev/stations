package main


import (
	"os"
)


func main() {
	mapfile, outputfile, ok := ParseArgs()
	if !ok {
		Log("Program exited with errors.")
		return
	}

	contents := ReadMapFile(mapfile)

	stations, connections, ok := ParseMap(string(contents))
	if !ok {
		Log("Program exited with errors.")
		return
	}

	dot := ToDot(stations, connections)

	var err error
	err = os.WriteFile(outputfile, []byte(dot), 0666)
	if err != nil {
		PrintErrArgs("os.WriteFile:", err.Error())
	} else {
		Log("Wrote output to", outputfile)
	}	

}
