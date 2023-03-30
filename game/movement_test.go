package game

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"math/rand"
)

func Test_SpawnPip(t *testing.T) {
	rand := rand.New(rand.NewSource(99))
	board := Board{[]int{}, rand}
	var result = SpawnPip(board)
	assert.Equal(t, board, result)
}
