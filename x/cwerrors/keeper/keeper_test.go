package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	e2eTesting "github.com/dymensionxyz/rollapp-wasm/e2e/testing"
)

type KeeperTestSuite struct {
	suite.Suite

	chain *e2eTesting.TestChain
}

func (s *KeeperTestSuite) SetupTest() {
	s.chain = e2eTesting.NewTestChain(s.T(), 1)
}

func TestCWErrorsKeeper(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
