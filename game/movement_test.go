package game

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"math/rand"
	"pips/types"
)

func Test_SpawnPip(t *testing.T) {
	rand := rand.New(rand.NewSource(99))
	matrix := []int{}
	board := types.Board{Matrix: matrix, Rand: rand}
	var result = SpawnPip(board)
	assert.Equal(t, board, result)
}

func Test_GetSpawningPoint(T *testing.T) {
	rand := rand.New(rand.NewSource(99))
	result := getSpawningPoint(rand)
	assert.Equal(0, result)
}
