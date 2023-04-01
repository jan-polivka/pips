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
	board := types.Board{Matrix: matrix, Cols: 10, Rand: rand}

	var result = SpawnPip(board)
	expectedIdx := 2
	assert.Equal(t, 1, result.Matrix[expectedIdx])
	assert.Equal(t, 1, len(result.Pips))
}

func Test_GetSpawningPoint(t *testing.T) {
	rand := rand.New(rand.NewSource(99))

	result := getSpawningPoint(rand)
	expected := 2
	assert.Equal(t, expected, result)
}

func Test_GetNextPosition(t *testing.T) {
	rand := rand.New(rand.NewSource(99))
	result := getNextPosition(types.Pip{Position: 1, Team: 1}, rand)
	expected := 2
	assert.Equal(t, expected, result)
}

func Test_CheckIfNewPositionIsValidBelowZero(t *testing.T) {
	rand := rand.New(rand.NewSource(99))
	matrix := [10]int{}
	board := types.Board{Matrix: matrix, Cols: 10, Rand: rand}
	result := isNewPositionValid(-1, board)
	assert.False(t, result)
}

func Test_CheckIfNewPositionIsValidAboveMaxColumns(t *testing.T) {
	rand := rand.New(rand.NewSource(99))
	const cols = 10
	matrix := [cols]int{}
	board := types.Board{Matrix: matrix, Cols: cols, Rand: rand}
	result := isNewPositionValid(10, board)
	assert.False(t, result)
}

func Test_isNewPositionValidOccupied(t *testing.T) {
	rand := rand.New(rand.NewSource(99))
	const cols = 10
	matrix := [cols]int{}
	const newPosition = 2
	matrix[newPosition] = 1
	board := types.Board{Matrix: matrix, Cols: cols, Rand: rand}
	result := isNewPositionValid(newPosition, board)
	assert.False(t, result)
}

func Test_isNewPositionValid(t *testing.T) {
	rand := rand.New(rand.NewSource(99))
	const cols = 10
	matrix := [cols]int{}
	board := types.Board{Matrix: matrix, Cols: cols, Rand: rand}
	result := isNewPositionValid(3, board)
	assert.True(t, result)
}

func Test_processNewPosition(t *testing.T) {
	newPosition := 1
	rand := rand.New(rand.NewSource(99))
	const cols = 10

	matrix := [cols]int{}
	matrix[2] = 1
	pip := types.Pip{Position: 2, Team: 1}
	pipArray := []types.Pip{pip}
	board := types.Board{Matrix: matrix, Cols: cols, Pips: pipArray, Rand: rand}
	result := processNewPosition(newPosition, board, 0)

	expectedMatrix := [cols]int{}
	expectedMatrix[newPosition] = 1
	expectedBoard := types.Board{Matrix: expectedMatrix, Cols: cols, Rand: rand}

	assert.Equal(t, expectedBoard.Matrix, result.Matrix)
	assert.Equal(t, newPosition, board.Pips[0].Position)
}
