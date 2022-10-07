package suite

import (
	"my-demo-service/app"
	"net"

	"github.com/stretchr/testify/suite"
)

// APISuite serve API feature tests
type APISuite struct {
	suite.Suite

	App      *app.Application
	Listener net.Listener
}
