package main

import (
	"math/rand"
	"pips/game"
	"pips/types"
)

func main() {

	rand := rand.New(rand.NewSource(99))
	matrix := []int{}
	board := types.Board{Matrix: matrix, Rand: rand}
	game.SpawnPip(board)

}
