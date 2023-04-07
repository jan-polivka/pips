package game

import (
	"math/rand"
	"pips/types"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	board types.Board
}

func (suite *TestSuite) SetupTest() {
	rand := rand.New(rand.NewSource(99))
	matrix := [10]int{}
	suite.board = types.Board{Matrix: matrix, Cols: 10, Rand: rand}
}

func (suite *TestSuite) TestExample() {
}

func Test_thisSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
