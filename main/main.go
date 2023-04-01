package main

import (
	"math/rand"
	"pips/game"
	"pips/random"
	"pips/types"
)

func main() {

	rand := rand.New(rand.NewSource(99))
	matrix := [10]int{}
	board := types.Board{Matrix: matrix, Rand: rand}
	game.SpawnPip(board, random.GenerateRandomInt)

}
