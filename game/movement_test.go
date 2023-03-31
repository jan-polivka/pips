package game

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"math/rand"
	"pips/types"
)

func Test_SpawnPip(t *testing.T) {
	rand := rand.New(rand.NewSource(99))
	matrix := [10]int{}
	board := types.Board{Matrix: matrix, Rand: rand}
	var result = SpawnPip(board)
	expectedIdx := 2
	assert.Equal(t, 1, result.Matrix[expectedIdx])
}

func Test_GetSpawningPoint(t *testing.T) {
	rand := rand.New(rand.NewSource(99))

	result := getSpawningPoint(rand)
	expected := 2
	assert.Equal(t, expected, result)
}
