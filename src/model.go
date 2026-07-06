package main


type Station struct {
	Name	string
	X		int
	Y		int
}

// TODO: Stations could be stored as actual Station struct instead of name string
type Connection struct {
	Begin	string
	End		string
}
