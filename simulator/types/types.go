package types

import (
	"math/rand"
)

type Board struct {
	History [][10]int
	Matrix  [10]int
	Cols    int
	Pips    []Pip
	Rand    *rand.Rand
}

type Pip struct {
	Position int
	Team     int
}
