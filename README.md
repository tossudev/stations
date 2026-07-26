![Stations Pathfinder](logo.gif)

A command line tool that finds the most efficient paths to move trains from one destination to another.  

## Prerequisites

- Go 1.25.x

## Features
### Base
- Parses user input map and creates graph from it
- Runs a max flow algorithm (Edmonds-Karp) to get max flow
- Finds paths from max flow
- Uses paths to create a train schedule and prints it out

### Extra
- Enhanced error handling
- Test suite

### Bonus
- Verbose mode (prints detailed output)

## Installation & Usage

1. Clone the repository  
```
git clone https://github.com/tossudev/stations  
cd stations  
```

2. Run the program  
```
go run ./cmd [MAP FILE] [START STATION] [TARGET STATION] [TRAINS AMOUNT] [optional args...]  
```
Optional arguments:  
- -v, --verbose  
  Verbose output  
- -h, --help  
  Displays usage  

Example:  
```
$ go run ./cmd data/london.map waterloo st_pancras 4
T1-victoria T2-euston
T1-st_pancras T2-st_pancras T3-victoria T4-euston
T3-st_pancras T4-st_pancras
```

## Testing

To run all test cases, run:  
```
go test ./cmd  
```

To run individual tests, run:  
```
go test ./cmd -run [TEST]  
```

The predefined tests are:  
	- London  
	- Fantasy  
	- Biomes  
	- Numbers  
	- Numbers2  
	- Classical  
	- Terminus  

To read about the test requirements, see [main_test.go](cmd/main_test.go)  
  
### Chickens
🪿 🪿 🪿

