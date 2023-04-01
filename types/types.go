package types

import (
	"math/rand"
)

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
