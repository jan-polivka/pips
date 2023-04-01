package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GenerateRandomInt(t *testing.T) {
	const min = 0
	const max = 100
	const size = max - min
	randMap := make(map[int]int)
	for i := 0; i < 100*size; i++ {
		result := GenerateRandomInt(min, max)
		randMap[result] = 1
	}

	for i := min; i < max; i++ {
		assert.NotEqual(t, 0, randMap[i])
	}
}