package types

import (
	"math/rand"
)

type Board struct {
	Matrix [10]int
	Rand   *rand.Rand
}
