package main


import (
	"os"
)


func main() {
	if len(os.Args) != 3 {
		PrintErr(ErrArgsCount)
		return
	}

	input := os.Args[1]
	output := os.Args[2]
	
	contents, err := os.ReadFile(input)
	if err != nil {
		PrintErr(err.Error())
		return
	}
	
	Log("Using", input, "as input.")

	stations, connections, ok := ParseMap(string(contents))
	if !ok {
		Log("Program exited with errors.")
		return
	}

	dot := ToDot(stations, connections)

	err = os.WriteFile(output, []byte(dot), 0666)
	if err != nil {
		PrintErr(err.Error())
		return
	}
	Log("Wrote output to", output)
	Log("Program exited successfully.")
}
