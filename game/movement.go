package game

import (
	"math/rand"
	"pips/types"
)

func SpawnPip(board types.Board) types.Board {
	spawningPoint := getSpawningPoint(board.Rand)
	board.Matrix[spawningPoint] = 1
	// pip := types.Pip{Position: spawningPoint, Team: 1}
	return board
}

func getSpawningPoint(rand *rand.Rand) int {
	return rand.Intn(5-0) + 0
}
