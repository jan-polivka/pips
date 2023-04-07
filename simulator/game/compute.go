package game

import (
	"encoding/json"
	"fmt"

	"pips/random"
	"pips/types"
)

func ComputeMatch(board types.Board) types.Board {
	head := board
	head.History = make([][10]int, 0)
	newBoard := MovePip(board, 0, random.GenerateRandomInt)
	head.History = append(head.History, newBoard.Matrix)
	var idx int
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			idx = 0
		} else {
			idx = 1
		}
		if newBoard.Pips[idx].Position != -1 {
			fmt.Println(newBoard.Pips[idx])
			newBoard = MovePip(newBoard, idx, random.GenerateRandomInt)
			head.History = append(head.History, newBoard.Matrix)
		} else {
			break
		}
	}
	return head
}

func MarshalMatch(board types.Board) string {
	res, err := json.Marshal(board.History)
	if err != nil {
		fmt.Println("Something went wrong")
	}
	return string(res)
}
