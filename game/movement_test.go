package game

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"fmt"
	"math/rand"
	"pips/types"
)

func Test_SpawnPip(t *testing.T) {
	rand := rand.New(rand.NewSource(99))
	matrix := [10]int{}
	board := types.Board{Matrix: matrix, Rand: rand}
	var result = SpawnPip(board)
	// expectedIdx := rand.Intn(5-0) + 0
	// assert.Equal(t, 1, result.Matrix[expectedIdx])
	assert.Equal(t, board, result)
}

func Test_GetSpawningPoint(t *testing.T) {
	rand := rand.New(rand.NewSource(99))

	result := getSpawningPoint(rand)
	expected := 2
	fmt.Println(result)
	fmt.Println(rand.Intn(5))
	fmt.Println(rand.Intn(5))
	fmt.Println(rand.Intn(5))
	fmt.Println(rand.Intn(5))
	assert.Equal(t, expected, result)
}
