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

func movePip(board types.Board, pipIdx int) types.Board {
	newPosition := getNextPosition(board.Pips[pipIdx], board.Rand)
	if isNewPositionValid(newPosition, board) {
		processNewPosition(newPosition, board, pipIdx)
	}
	return board
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

func processNewPosition(newPosition int, board types.Board, pipIdx int) types.Board {
	board.Matrix[board.Pips[pipIdx].Position] = 0
	board.Pips[pipIdx].Position = newPosition
	board.Matrix[newPosition] = board.Pips[pipIdx].Team
	return board
}
