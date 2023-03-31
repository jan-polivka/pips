package game

import (
	"fmt"
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

func Test_GetSpawningPoint(t *testing.T) {
	rand := rand.New(rand.NewSource(99))

	randNumber := rand.Int()
	fmt.Println(randNumber)

	result := getSpawningPoint(rand)
	expected := 1
	assert.Equal(t, expected, result)
}
