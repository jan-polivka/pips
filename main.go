package main

import (
	"fmt"
	"math/rand"
	"pips/game"
	"pips/random"
	"pips/types"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	const cols = 10
	matrix := [cols]int{}
	board := types.Board{Matrix: matrix, Cols: cols, Pips: []types.Pip{}, Rand: rand}
	board = game.SpawnPip(board, random.GenerateRandomInt)
	fmt.Println(board.Matrix)
	board = game.MovePip(board, 0, random.GenerateRandomInt)
	fmt.Println(board.Matrix)
	board = game.MovePip(board, 0, random.GenerateRandomInt)
	fmt.Println(board.Matrix)
	board = game.MovePip(board, 0, random.GenerateRandomInt)
	fmt.Println(board.Matrix)
	board = game.MovePip(board, 0, random.GenerateRandomInt)
	fmt.Println(board.Matrix)
	board = game.MovePip(board, 0, random.GenerateRandomInt)
	fmt.Println(board.Matrix)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.Run()
}
