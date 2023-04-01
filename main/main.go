package Main

import (
	"fmt"
	"math/rand"
	"pips/game"
	"pips/random"
	"pips/types"
)

func main() {

	rand := rand.New(rand.NewSource(99))
	matrix := [10]int{}
	board := types.Board{Matrix: matrix, Rand: rand}
	board = game.SpawnPip(board, random.GenerateRandomInt)
	board = game.MovePip(board, 0, random.GenerateRandomInt)
	fmt.Println(board)
	board = game.MovePip(board, 0, random.GenerateRandomInt)
	fmt.Println(board)
	board = game.MovePip(board, 0, random.GenerateRandomInt)
	fmt.Println(board)
	board = game.MovePip(board, 0, random.GenerateRandomInt)
	fmt.Println(board)
	board = game.MovePip(board, 0, random.GenerateRandomInt)
	fmt.Println(board)
}
