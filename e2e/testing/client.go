package e2eTesting

import (
	"context"
	"fmt"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"google.golang.org/grpc"

	"github.com/dymensionxyz/rollapp-wasm/app"
)

var _ grpc.ClientConnInterface = (*grpcClient)(nil)

type grpcClient struct {
	app *app.App
}

func (c grpcClient) Invoke(ctx context.Context, method string, args, reply interface{}, opts ...grpc.CallOption) error {
	req := args.(codec.ProtoMarshaler)
	resp := c.app.Query(abci.RequestQuery{
		Data:   c.app.AppCodec().MustMarshal(req),
		Path:   method,
		Height: 0, // TODO: heightened queries
		Prove:  false,
	})

	if resp.Code != abci.CodeTypeOK {
		return fmt.Errorf(resp.Log)
	}

	c.app.AppCodec().MustUnmarshal(resp.Value, reply.(codec.ProtoMarshaler))

	return nil
}

func (c grpcClient) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	panic("not supported")
}

func (chain *TestChain) Client() grpc.ClientConnInterface {
	return grpcClient{app: chain.app}
}
