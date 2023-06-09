package game

import (
	"fmt"
	"math/rand"
	"pips/types"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	board    types.Board
	posOne   int
	posTwo   int
	teamOne  int
	teamTwo  int
	randFunc []func(int, int) int
}

func (suite *TestSuite) SetupTest() {
	rand := rand.New(rand.NewSource(99))
	matrix := [10]int{}
	suite.board = types.Board{Matrix: matrix, Cols: 10, Rand: rand}
	suite.posOne = 2
	suite.posTwo = 8
	suite.teamOne = 1
	suite.teamTwo = 2
	suite.board = SpawnPip(suite.board, func(int, int) int { return suite.posOne }, suite.teamOne)
	suite.board = SpawnPip(suite.board, func(int, int) int { return suite.posTwo }, suite.teamTwo)
	suite.randFunc = []func(int, int) int{func(int, int) int { return 1 }, func(i1, i2 int) int { return -1 }}
}

func (suite *TestSuite) Test_moveTwoPips() {
	newBoard := movePipInTest(suite, suite.board, 0)
	newBoard = movePipInTest(suite, newBoard, 1)
	expectedBoard := [10]int{0, 0, 0, 1, 0, 0, 0, 2, 0, 0}
	suite.Equal(expectedBoard, newBoard.Matrix)
}

func (suite *TestSuite) Test_pipOneEatsPipTwo() {
	newBoard := ProcessNewPosition(suite.posTwo-1, suite.board, 0)
	newBoard = movePipInTest(suite, newBoard, 0)
	expectedBoard := [10]int{0, 0, 0, 0, 0, 0, 0, 0, 1, 0}
	expectedBoardPip := suite.board.Pips[1]
	expectedBoardPip.Position = -1
	suite.Equal(expectedBoard, newBoard.Matrix)
	suite.Equal(expectedBoardPip, newBoard.Pips[1])
}

func (suite *TestSuite) Test_eatPipIfThreatenedTrue() {
	result := eatPipIfThreatened(suite.board, suite.posTwo, 0)
	expected := suite.board.Pips[1]
	expected.Position = -1
	suite.Equal(expected, result.Pips[1])
}

func (suite *TestSuite) Test_eatPipIfThreatenedFalse() {
	result := eatPipIfThreatened(suite.board, suite.posTwo, 0)
	expected := suite.board.Pips[0]
	suite.Equal(expected, result.Pips[0])
}

func (suite *TestSuite) Test_eatPipIfThreatenedSelfFail() {
	expected := suite.board.Pips[0]
	fmt.Println(expected)
	result := eatPipIfThreatened(suite.board, suite.posOne, 0)
	suite.Equal(expected, result.Pips[0])
}

func Test_thisSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func movePipInTest(testSuite *TestSuite, newBoard types.Board, currPip int) types.Board {
	return MovePip(newBoard, currPip, testSuite.randFunc[currPip])
}
