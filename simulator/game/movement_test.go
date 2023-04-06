package game

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"math/rand"
	"pips/types"
)

func Test_SpawnPip(t *testing.T) {
	rand := rand.New(rand.NewSource(99))
	matrix := [10]int{}
	board := types.Board{Matrix: matrix, Cols: 10, Rand: rand}

	var result = SpawnPip(board, func(int, int) int { return 4 }, 1)
	expectedIdx := 4
	assert.Equal(t, 1, result.Matrix[expectedIdx])
	assert.Equal(t, 1, len(result.Pips))
}

func Test_SpawnTwoPips(t *testing.T) {
	rand := rand.New(rand.NewSource(99))
	matrix := [10]int{}
	board := types.Board{Matrix: matrix, Cols: 10, Rand: rand}

	const posOne = 2
	const posTwo = 8
	const teamOne = 1
	const teamTwo = 2

	var result = SpawnPip(board, func(int, int) int { return posOne }, teamOne)
	result = SpawnPip(result, func(int, int) int { return posTwo }, teamTwo)
	fmt.Println(result)
	assert.Equal(t, 1, result.Matrix[2])
	assert.Equal(t, 2, result.Matrix[8])
}

func Test_GetNextPosition(t *testing.T) {
	result := getNextPosition(types.Pip{Position: 1, Team: 1}, func(int, int) int { return 1 })
	expected := 2
	assert.Equal(t, expected, result)
}

func Test_isNewPositionValidBelowZero(t *testing.T) {
	rand := rand.New(rand.NewSource(99))
	matrix := [10]int{}
	board := types.Board{Matrix: matrix, Cols: 10, Rand: rand}
	result := isNewPositionValid(-1, board)
	assert.False(t, result)
}

func Test_isNewPositionValidAboveMaxColumns(t *testing.T) {
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

func Test_spawnAndMove(t *testing.T) {

	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	const cols = 10
	matrix := [cols]int{}
	board := types.Board{Matrix: matrix, Cols: cols, Pips: []types.Pip{}, Rand: rand}
	randFunc := func(int, int) int { return 2 }
	board = SpawnPip(board, randFunc, 1)
	randFunc = func(int, int) int { return 1 }
	newBoard := MovePip(board, 0, randFunc)
	for i := 0; i < 6; i++ {
		newBoard = MovePip(newBoard, 0, randFunc)
	}
	expectedMatrix := [cols]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	assert.Equal(t, expectedMatrix, newBoard.Matrix)

}

func Test_GetBoundsTeamOne(t *testing.T) {
	l, u := getBounds(1)
	assert.Equal(t, 0, l)
	assert.Equal(t, 3, u)
}
