package internal

import (
	"fmt"
	"strings"
)

// TODO: take into account small number of trains
// this approach is primitive in the sense that it uses only the maxflow paths
// there are cases where it would be more efficient to find shorter paths
// however, this is trivial for the project requirements
func CreateSchedule(graphList *GraphList, paths [][]int, source, sink string, numOfTrains int, printResult bool) int {
	for i, path := range paths {
		output := fmt.Sprintf("Path #%d: ", i+1)

		for _, station := range path {
			output += fmt.Sprintf("%s -> ", graphList.stationsNames[station%len(graphList.stationsNames)])
		}
		output = output[:len(output)-3]
		Log(output)
	}	

	type Train struct {
		id       int
		path     []int
		position int
	}

	assignments := assignTrains(paths, numOfTrains)

	var active []*Train
	finished := 0
	turn := 0

	for finished < numOfTrains {
		turn++
		var output []string

		var stillActive []*Train
		for _, train := range active {
			train.position++

			if train.position < len(train.path) {
				station := graphList.stationsNames[train.path[train.position]]
				output = append(output, fmt.Sprintf("T%d-%s", train.id, station))
				stillActive = append(stillActive, train)
			} else {
				finished++
			}
		}
		active = stillActive

		// send new trains, one per turn
		for pathIdx, trainList := range assignments {
			if len(trainList) == 0 {
				continue
			}

			trainID := trainList[0]
			assignments[pathIdx] = trainList[1:]

			path := paths[pathIdx]
			// start at first skipping source
			if len(path) > 1 {
				train := &Train{
					id:       trainID,
					path:     path,
					position: 1,
				}
				station := graphList.stationsNames[path[1]]
				output = append(output, fmt.Sprintf("T%d-%s", train.id, station))

				active = append(active, train)

			}
		}
		if len(output) > 0 && printResult {
			fmt.Println(strings.Join(output, " "))
		}
	}

	return turn - 1
}

func assignTrains(paths [][]int, numTrains int) [][]int {
	type PathInfo struct {
		index int
		load  int
	}

	infos := make([]PathInfo, len(paths))
	for i := range paths {
		infos[i] = PathInfo{
			index: i,
			load:  len(paths[i]), // initial cost is path length
		}
	}

	assignments := make([][]int, len(paths))

	for trainID := 1; trainID <= numTrains; trainID++ {
		best := 0
		for i := 1; i < len(infos); i++ {
			if infos[i].load < infos[best].load {
				best = i
			}
		}

		pathIdx := infos[best].index
		assignments[pathIdx] = append(assignments[pathIdx], trainID)

		infos[best].load++ // path becomes busier
	}
	return assignments
}
