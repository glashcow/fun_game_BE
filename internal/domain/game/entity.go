package game

import "time"

type Game struct {
	ID         string
	PlayerIDs  []string
	BingoCards []*BingoCard
	Calls      []BingoCalls
}

type BingoCard struct {
	ID       string
	PlayerID string
	Numbers  map[Coordinate]Number
}

type BingoCalls struct {
	ID       string
	Number   int
	CalledAt time.Time
}
