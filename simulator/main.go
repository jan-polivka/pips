package main

import (
	"fmt"
	"math/rand"
	"pips/game"
	"pips/random"
	"pips/types"
	"time"

	"github.com/gin-contrib/cors"
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
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET"},
		AllowHeaders: []string{"Origin"},
	}))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.GET("/match", func(c *gin.Context) {
		fmt.Println("here in my garage")
		c.JSON(200, gin.H{"message": "sent"})
	})
	r.Run()
}
