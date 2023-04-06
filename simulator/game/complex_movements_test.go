package game

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	myVariable int
}

func (suite *TestSuite) SetupTest() {
	suite.myVariable = 5
}

func (suite *TestSuite) TestExample() {
	suite.Equal(5, suite.myVariable)
}

func Test_thisSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
