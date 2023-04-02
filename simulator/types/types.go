package types

import (
	"math/rand"
)

type Board struct {
	History BoardSnapshot
	Matrix  [10]int
	Cols    int
	Pips    []Pip
	Rand    *rand.Rand
}

type Pip struct {
	Position int
	Team     int
}

type BoardSnapshot struct {
	Board [10]int
	Next  *BoardSnapshot
}
