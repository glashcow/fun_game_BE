package game

type Number struct {
	Value   int
	checked bool
}

// Coordinate (0, 0) [(x, y)] would correspond to the top left corner, with x values increasing to the left (max 7).
// and y increasing down to the bottom (max 2).
type Coordinate struct {
	x int
	y int
}
