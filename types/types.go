package types

import (
	"math/rand"
)

type Board struct {
	Matrix []int
	Rand   *rand.Rand
}
