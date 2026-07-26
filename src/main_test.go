package main

import (
	"testing"
)

type Suite struct {
	// Input
	File		string
	Start		string
	End			string
	Trains		int
	// Expected output
	Maxflow		int
	Turns		int
}

const (
	TestDirectory	string = "../data/"
)

// TODO: Does not test for turns yet, only maxflow
func testAny(suite Suite, t *testing.T) {
	contents := ReadMapFile(TestDirectory + suite.File)
	graphList, _ := ParseMap(string(contents), suite.Start, suite.End)
	maxflow, paths := MaxFlow(graphList, suite.Start, suite.End)

	if maxflow != suite.Maxflow {
		t.Errorf(`Test failed: %s. Maxflow should be %d, is %d.`, suite.File, suite.Maxflow, maxflow)
	}

	turns := CreateSchedule(graphList, paths, suite.Start, suite.End, suite.Trains, false)
	if turns > suite.Turns {
		t.Errorf(`Test failed: %s. Should take %d turns, takes %d.`, suite.File, suite.Turns, turns)
	}
}

// It can find more than one route for 2 trains between waterloo and st_pancras for the London Network Map:
func TestLondon(t *testing.T) {
	suite := Suite{
		File:		"london.map",
		Start:		"waterloo",
		End:		"euston",
		Trains:		2,
		Maxflow:	2,
		Turns:		6,
	}

	testAny(suite, t)
}

// It completes the movements in no more than 6 turns for 4 trains between bond_square and space_port:
func TestFantasy(t *testing.T) {
	suite := Suite{
		File:		"fantasy.map",
		Start:		"bond_square",
		End:		"space_port",
		Trains:		4,
		Maxflow:	1,
		Turns:		6,
	}

	testAny(suite, t)
}

// It completes the movements in no more than 8 turns for 10 trains between jungle and desert:
func TestBiomes(t *testing.T) {
	suite := Suite{
		File:		"biomes.map",
		Start:		"jungle",
		End:		"desert",
		Trains:		10,
		Maxflow:	3,
		Turns:		8,
	}

	testAny(suite, t)
}

// It completes the movements in no more than 8 turns for 9 trains between small and large:
func TestNumbers(t *testing.T) {
	suite := Suite{
		File:		"numbers.map",
		Start:		"small",
		End:		"large",
		Trains:		9,
		Maxflow:	4,
		Turns:		8,
	}

	testAny(suite, t)
}

// It completes the movements in no more than 6 turns for 4 trains between two and four:
func TestNumbers2(t *testing.T) {
	suite := Suite{
		File:		"numbers2.map",
		Start:		"two",
		End:		"four",
		Trains:		4,
		Maxflow:	1,
		Turns:		6,
	}

	testAny(suite, t)
}

// It completes the movements in no more than 6 turns for 9 trains between beethoven and part:
func TestClassical(t *testing.T) {
	suite := Suite{
		File:		"classical.map",
		Start:		"beethoven",
		End:		"part",
		Trains:		9,
		Maxflow:	2,
		Turns:		6,
	}

	testAny(suite, t)
}

// It completes the movements in no more than 11 turns for 20 trains between beginning and terminus:
func TestTerminus(t *testing.T) {
	suite := Suite{
		File:		"terminus.map",
		Start:		"beginning",
		End:		"terminus",
		Trains:		20,
		Maxflow:	2,
		Turns:		11,
	}

	testAny(suite, t)
}
