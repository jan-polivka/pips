package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_spawnPip(t *testing.T) {
	board := []int{}
	var result = spawnPip(board)
	// var wrong [5]int
	assert.Equal(t, board, result)
}
