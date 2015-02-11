package app

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type AppSuite struct {
	suite.Suite
}

func TestAppSuite(t *testing.T) {
	suite.Run(t, new(AppSuite))
}

func (s *AppSuite) SetupTest() {
	// Test Setup
}

func (s *AppSuite) SetupSuite() {
	// Suite Setup
}

func (s *AppSuite) TestTitleGet() {
	// A test
}
