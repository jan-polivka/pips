package types

import (
	"math/rand"
)

// I should remove the Matrix and replace it with x and y limits
// Then add an array of pips
type Board struct {
	Matrix [10]int
	Rand   *rand.Rand
}

type Pip struct {
	Position [2]int
	Team     int
}
