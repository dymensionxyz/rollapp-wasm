package keeper_test

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	e2eTesting "github.com/dymensionxyz/rollapp-wasm/e2e/testing"
	"github.com/dymensionxyz/rollapp-wasm/pkg/testutils"
	cwerrorsKeeper "github.com/dymensionxyz/rollapp-wasm/x/cwerrors/keeper"
	"github.com/dymensionxyz/rollapp-wasm/x/cwerrors/types"
)

func (s *KeeperTestSuite) TestSubscribeToError() {
	// Setting up chain and contract in mock wasm keeper
	ctx, keeper := s.chain.GetContext(), s.chain.GetApp().CWErrorsKeeper
	contractViewer := testutils.NewMockContractViewer()
	keeper.SetWasmKeeper(contractViewer)
	contractAddr := e2eTesting.GenContractAddresses(2)[0]
	contractAddr2 := e2eTesting.GenContractAddresses(2)[1]
	contractAdminAcc := s.chain.GetAccount(2)
	contractViewer.AddContractAdmin(
		contractAddr.String(),
		contractAdminAcc.Address.String(),
	)
	params, err := keeper.GetParams(ctx)
	s.Require().NoError(err)
	params.SubscriptionFee = sdk.NewInt64Coin(sdk.DefaultBondDenom, 1)
	err = keeper.SetParams(ctx, params)
	s.Require().NoError(err)

	expectedEndHeight := ctx.BlockHeight() + params.SubscriptionPeriod

	msgServer := cwerrorsKeeper.NewMsgServer(keeper)

	testCases := []struct {
		testCase    string
		input       func() *types.MsgSubscribeToError
		expectError bool
		errorType   error
	}{
		{
			testCase: "FAIL: empty request",
			input: func() *types.MsgSubscribeToError {
				return nil
			},
			expectError: true,
			errorType:   status.Error(codes.InvalidArgument, "empty request"),
		},
		{
			testCase: "FAIL: invalid sender address",
			input: func() *types.MsgSubscribeToError {
				return &types.MsgSubscribeToError{
					Sender:          "👻",
					ContractAddress: contractAddr.String(),
					Fee:             sdk.NewInt64Coin(sdk.DefaultBondDenom, 100),
				}
			},
			expectError: true,
			errorType:   errors.New("invalid bech32 string length 4"),
		},
		{
			testCase: "FAIL: invalid contract address",
			input: func() *types.MsgSubscribeToError {
				return &types.MsgSubscribeToError{
					Sender:          contractAdminAcc.Address.String(),
					ContractAddress: "👻",
					Fee:             sdk.NewInt64Coin(sdk.DefaultBondDenom, 100),
				}
			},
			expectError: true,
			errorType:   errors.New("invalid bech32 string length 4"),
		},
		{
			testCase: "FAIL: contract not found",
			input: func() *types.MsgSubscribeToError {
				return &types.MsgSubscribeToError{
					Sender:          contractAdminAcc.Address.String(),
					ContractAddress: contractAddr2.String(),
					Fee:             sdk.NewInt64Coin(sdk.DefaultBondDenom, 100),
				}
			},
			expectError: true,
			errorType:   types.ErrContractNotFound,
		},
		{
			testCase: "FAIL: account doesnt have enough balance",
			input: func() *types.MsgSubscribeToError {
				return &types.MsgSubscribeToError{
					Sender:          contractAddr.String(),
					ContractAddress: contractAddr.String(),
					Fee:             params.SubscriptionFee,
				}
			},
			expectError: true,
			errorType:   sdkerrors.ErrInsufficientFunds,
		},
		{
			testCase: "OK: valid request",
			input: func() *types.MsgSubscribeToError {
				return &types.MsgSubscribeToError{
					Sender:          contractAdminAcc.Address.String(),
					ContractAddress: contractAddr.String(),
					Fee:             params.SubscriptionFee,
				}
			},
			expectError: false,
		},
	}
	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case: %s", tc.testCase), func() {
			req := tc.input()
			res, err := msgServer.SubscribeToError(sdk.WrapSDKContext(ctx), req)
			if tc.expectError {
				s.Require().Error(err)
				s.Assert().ErrorContains(err, tc.errorType.Error())
			} else {
				s.Require().NoError(err)
				s.Require().Equal(expectedEndHeight, res.SubscriptionValidTill)
			}
		})
	}
}

