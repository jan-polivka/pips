package game

import (
	"math/rand"
	"pips/types"
)

func SpawnPip(board types.Board) types.Board {
	spawningPoint := getSpawningPoint(board.Rand)
	board.Matrix[spawningPoint] = 1
	pip := types.Pip{Position: spawningPoint, Team: 1}
	board.Pips = append(board.Pips, pip)
	return board
}

func getSpawningPoint(rand *rand.Rand) int {
	return rand.Intn(5-0) + 0
}

func movePip(board types.Board) {
	//get a new position
	//check if new position is taken
	//if not
	//delete old pip in the board

}
