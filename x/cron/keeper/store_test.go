package keeper_test

import (
	cronTypes "github.com/dymensionxyz/rollapp-wasm/x/cron/types"
)

func (s *KeeperTestSuite) TestSetContract() {

	cronKeeper, ctx := &s.app.CronKeeper, &s.ctx
	params := cronKeeper.GetParams(*ctx)

	for _, tc := range []struct {
		name   string
		msg    cronTypes.WhitelistedContract
		ExpErr error
	}{
		{
			"Single Player Registered Contract successfully",
			cronTypes.WhitelistedContract{
				GameId:          1,
				SecurityAddress: params.GetSecurityAddress()[0],
				ContractAdmin:   s.addr(2).String(),
				GameName:        "coinflip",
				ContractAddress: s.addr(3).String(),
				GameType:        1,
			},
			nil,
		},
		{
			"Multi Player Registered Contract successfully",
			cronTypes.WhitelistedContract{
				GameId:          2,
				SecurityAddress: params.GetSecurityAddress()[0],
				ContractAdmin:   s.addr(2).String(),
				GameName:        "roulette",
				ContractAddress: s.addr(4).String(),
				GameType:        2,
			},
			nil,
		},
	} {
		s.Run(tc.name, func() {
			err := cronKeeper.SetContract(*ctx, tc.msg)
			if tc.ExpErr != nil {
				s.Require().Error(err)
				s.Require().EqualError(err, tc.ExpErr.Error())
			} else {

				s.Require().NoError(err)
				res, found := cronKeeper.GetWhitelistedContract(*ctx, tc.msg.GameId)
				s.Require().True(found)
				s.Require().Equal(res.GameId, tc.msg.GameId)
				s.Require().Equal(res.SecurityAddress, tc.msg.SecurityAddress)
				s.Require().Equal(res.ContractAdmin, tc.msg.ContractAdmin)
				s.Require().Equal(res.GameName, tc.msg.GameName)
				s.Require().Equal(res.ContractAddress, tc.msg.ContractAddress)
				s.Require().Equal(res.GameType, tc.msg.GameType)
			}
		})
	}
}
