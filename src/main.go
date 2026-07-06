package main


import (
	"os"
	"fmt"
)


func main() {
	if len(os.Args) != 3 {
		fmt.Println("Give input and output files as arguments!")
		os.Exit(1)
	}

	input := os.Args[1]
	output := os.Args[2]
	
	fmt.Println("Using", input, "as input.")

	contents, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("os.ReadFile ERROR:", err)
	}

	var stations []Station
	var connections []Connection
	var parseErr error

	stations, connections, parseErr = ParseInput(string(contents))
	if parseErr != nil {
		fmt.Println("parseInput ERROR:", parseErr)
	}

	dot := ToDot(stations, connections)

	// For testing purposes only
	//fmt.Println(dot)

	err = os.WriteFile(output, []byte(dot), 0666)
	if err != nil {
		fmt.Println("os.WriteFile ERROR:", err)
	} else {
		fmt.Println("Wrote output to", output)
	}
}
