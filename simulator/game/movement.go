package game

import (
	"pips/random"
	"pips/types"
)

func SpawnPip(board types.Board, getRandomInt random.GenerateRandomIntFunction, team int) types.Board {
	lowerBound, upperBound := getBounds(team)
	spawningPoint := getRandomInt(lowerBound, upperBound)
	board.Matrix[spawningPoint] = team
	pip := types.Pip{Position: spawningPoint, Team: team}
	board.Pips = append(board.Pips, pip)
	return board
}

func MovePip(board types.Board, pipIdx int, getRandomInt random.GenerateRandomIntFunction) types.Board {
	newPosition := getNextPosition(board.Pips[pipIdx], getRandomInt)
	if isNewPositionValid(newPosition, board) {
		return ProcessNewPosition(newPosition, board, pipIdx)
	}
	return board
}

func getNextPosition(pip types.Pip, getRandomInt random.GenerateRandomIntFunction) int {
	newPosition := pip.Position + getRandomInt(-1, 2)
	return newPosition
}

func isNewPositionValid(newPosition int, board types.Board) bool {
	if newPosition < 0 || newPosition >= board.Cols {
		return false
	}
	return true
}

func ProcessNewPosition(newPosition int, board types.Board, pipIdx int) types.Board {
	//is a pip being eaten?
	// update the eaten pip
	board.Matrix[board.Pips[pipIdx].Position] = 0
	board.Pips[pipIdx].Position = newPosition
	board.Matrix[newPosition] = board.Pips[pipIdx].Team
	return board
}

func getBounds(team int) (int, int) {
	if team == 1 {
		return 0, 3
	} else if team == 2 {
		return 7, 9
	} else {
		panic("Wrong team numbering")
	}
}

func isAPipThreatened(board types.Board, position int) bool {
	return false
}
