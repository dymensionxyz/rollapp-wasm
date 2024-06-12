package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRegisterCron{}, "wasmrollapp/cron/MsgRegisterCron", nil)
	cdc.RegisterConcrete(&MsgUpdateCronJob{}, "wasmrollapp/cron/MsgUpdateCronJob", nil)
	cdc.RegisterConcrete(&MsgDeleteCronJob{}, "wasmrollapp/cron/MsgDeleteCronJob", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgRegisterCron{},
		&MsgUpdateCronJob{},
		&MsgDeleteCronJob{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	moduleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	RegisterLegacyAminoCodec(authzcodec.Amino)
	amino.Seal()
}
