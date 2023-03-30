package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_spawnPip(t *testing.T) {
	board := []int{1}
	var result = spawnPip(board)
	assert.Equal(t, board, result)
}
