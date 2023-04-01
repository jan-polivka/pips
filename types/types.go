package types

import (
	"math/rand"
)

// I should remove the Matrix and replace it with x and y limits
// Then add an array of pips
type Board struct {
	History BoardSnapshot
	Matrix  [10]int
	Cols    int
	Pips    [1]Pip
	Rand    *rand.Rand
}

type Pip struct {
	Position [2]int
	Team     int
}

type BoardSnapshot struct {
	Board [10]int
	Next  *BoardSnapshot
}
