package plugin

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PluginTestSuite struct {
	suite.Suite
}

func TestPluginTestSuite(t *testing.T) {
	suite.Run(t, new(PluginTestSuite))
}

//TODO write a couple of tests...
func (suite *PluginTestSuite) TestGetRegistries() {
	t := suite.T()
	assert.True(t, true)
}
