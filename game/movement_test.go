package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SpawnPip(t *testing.T) {
	board := []int{1}
	var result = SpawnPip(board)
	assert.Equal(t, board, result)
}
