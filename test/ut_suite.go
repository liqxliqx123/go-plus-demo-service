package test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

// Suite suite for UT
type Suite struct {
	suite.Suite
	Ctrl *gomock.Controller
}

// SetupTest implement suite.SetupTestSuite
func (t *Suite) SetupTest() {
	t.Ctrl = gomock.NewController(t.T())
}

// TearDownTest implement suite.TearDownTestSuite
func (t *Suite) TearDownTest() {
	t.Ctrl.Finish()
}
