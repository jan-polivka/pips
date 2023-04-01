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
	//assign the new position to the pip
	//assign the new position in the matrix
	//if yes
	//do nothing

}

func getNextPosition(pip types.Pip, rand *rand.Rand) int {
	newPosition := pip.Position + rand.Intn(2+1) - 1
	return newPosition
}

func isNewPositionValid(newPosition int, board types.Board) bool {
	if newPosition < 0 || newPosition >= board.Cols {
		return false
	}
	if board.Matrix[newPosition] != 0 {
		return false
	}
	return true
}

func isNewPositionWithinBounds(newPosition int, board types.Board) bool {
	if newPosition < 0 || newPosition >= board.Cols {
		return false
	}
	return true
}

func isNewPositionOccupied(newPosition int, board types.Board) bool {
	if board.Matrix[newPosition] != 0 {
		return false
	}
	return true
}
