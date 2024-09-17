package basic_demo

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type ExampleSuite struct {
	suite.Suite
}

func TestExampleSuite(t *testing.T) {
	suite.Run(t, &ExampleSuite{})
}

func (es *ExampleSuite) TestTrue() {
	es.T().Log("Running TestTrue...")
	es.True(true)
}

func (es *ExampleSuite) TestFalse() {
	es.T().Log("Running TestFalse...")
	es.False(false)
}

func (es *ExampleSuite) SetupSuite() {
	es.T().Log("Setting up the suite...")
}

func (es *ExampleSuite) TearDownSuite() {
	es.T().Log("Tearing down the suite...")
}

func (es *ExampleSuite) SetupTest() {
	es.T().Log("Setting up the test...")
}

func (es *ExampleSuite) TearDownTest() {
	es.T().Log("Tearing down the test...")
}

func (es *ExampleSuite) BeforeTest(suiteName, testName string) {
	es.T().Log("Before test: ", suiteName, testName)
}

func (es *ExampleSuite) AfterTest(suiteName, testName string) {
	es.T().Log("After test: ", suiteName, testName)
}
