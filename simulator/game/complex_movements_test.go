package game

import (
	"math/rand"
	"pips/types"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	board   types.Board
	posOne  int
	posTwo  int
	teamOne int
	teamTwo int
}

func (suite *TestSuite) SetupTest() {
	rand := rand.New(rand.NewSource(99))
	matrix := [10]int{}
	suite.board = types.Board{Matrix: matrix, Cols: 10, Rand: rand}
	suite.posOne = 2
	suite.posTwo = 8
	suite.teamOne = 1
	suite.teamTwo = 2
}

func (suite *TestSuite) TestExample() {
	var result = SpawnPip(suite.board, func(int, int) int { return suite.posOne }, suite.teamOne)
	result = SpawnPip(result, func(int, int) int { return suite.posTwo }, suite.teamTwo)
	randFunc := func(int, int) int { return 1 }
	newBoard := MovePip(result, 0, randFunc)
	var currPip int
	for i := 0; i < 3; i++ {
		if i%2 == 0 {
			currPip = 1
			randFunc = func(int, int) int { return -1 }
		} else {
			currPip = 0
			randFunc = func(int, int) int { return 1 }
		}
		newBoard = MovePip(newBoard, currPip, randFunc)
	}
	expectedBoard := [10]int{0, 0, 0, 0, 1, 0, 2, 0, 0, 0}
	suite.Equal(expectedBoard, newBoard.Matrix)

}

func Test_thisSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
