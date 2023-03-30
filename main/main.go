package main

import (
	"math/rand"
)

type Board struct {
	matrix []int
	rand   *rand.Rand
}

func main() {

	rand := rand.New(rand.NewSource(99))
	board := Board{[]int{}, rand}
	game.spawnPip(board)

}
