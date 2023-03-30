package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_spawnPip(t *testing.T) {
	var board [6]int
	var result = spawnPip(&board[0], 6)
	assert.Equal(t, &board[0], result)
}
