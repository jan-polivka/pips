package main

import (
	"math/rand"
	"pips/game"
	"pips/types"
)

func main() {

	rand := rand.New(rand.NewSource(99))
	board := types.Board{[]int{}, rand}
	game.SpawnPip(board)

}
