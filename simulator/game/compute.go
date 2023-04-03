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
	for i := 0; i < 100; i++ {
		newBoard := MovePip(board, 0, random.GenerateRandomInt)
		head.History = append(head.History, newBoard.Matrix)
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
