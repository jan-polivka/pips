package main

import (
	"math/rand"
	"pips/game"
	"pips/random"
	"pips/types"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

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
		c.JSON(200, gin.H{"message": game.MarshalMatch(doTheThing())})
	})
	r.Run()
}

func doTheThing() types.Board {

	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	const cols = 10
	matrix := [cols]int{}
	board := types.Board{Matrix: matrix, Cols: cols, Pips: []types.Pip{}, Rand: rand}
	board = game.SpawnPip(board, random.GenerateRandomInt, 1)
	board = game.SpawnPip(board, random.GenerateRandomInt, 2)
	board = game.ComputeMatch(board)
	return board
}
