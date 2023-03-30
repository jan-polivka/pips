package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_spawnPip(t *testing.T) {
	board := []int{}
	var result = spawnPip(board)
	wrong := []int{1}
	assert.Equal(t, wrong, result)
}
