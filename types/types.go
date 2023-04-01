package types

import (
	"math/rand"
)

// I should remove the Matrix and replace it with x and y limits
// Then add an array of pips
type Board struct {
	History [10][10]int
	Matrix  [10]int
	Pips    [1]Pip
	Rand    *rand.Rand
}

type Pip struct {
	Position [2]int
	Team     int
}
