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
	suite.Equal(expectedBoard, newBoard.Matrix)
}

func (suite *TestSuite) Test_isAPipThreatenedFalse() {
	result := isAPipThreatened(suite.board, suite.posOne+1)
	suite.Equal(false, result)
}

func (suite *TestSuite) Test_isAPipThreatenedTrue() {
	result := isAPipThreatened(suite.board, suite.posOne)
	suite.Equal(true, result)
}

func (suite *TestSuite) Test_eatPip() {
	result := eatPip(suite.board, suite.posOne)
	expected := suite.board.Pips[0]
	expected.Position = -1
	fmt.Println(result)
}

func Test_thisSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func movePipInTest(testSuite *TestSuite, newBoard types.Board, currPip int) types.Board {
	return MovePip(newBoard, currPip, testSuite.randFunc[currPip])
}
