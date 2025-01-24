package massalib

import (
	"context"

	"github.com/ModChain/massalib/massagrpc"
	"google.golang.org/grpc"
)

type RPC struct {
	cPub *grpc.ClientConn
	pub  massagrpc.PublicServiceClient
}

// New returns a new connection to a massa node
//
// grpc://localhost:33037
func New(target string, opts ...grpc.DialOption) (*RPC, error) {
	cPub, err := grpc.NewClient(target, opts...)
	if err != nil {
		return nil, err
	}
	cl := &RPC{
		cPub: cPub,
		pub:  massagrpc.NewPublicServiceClient(cPub),
	}

	return cl, nil
}

// Public exposes the raw grpc public interface
func (rpc *RPC) Public() massagrpc.PublicServiceClient {
	return rpc.pub
}

// GetStatus returns the status of the node
func (rpc *RPC) GetStatus(ctx context.Context) (*massagrpc.PublicStatus, error) {
	res, err := rpc.pub.GetStatus(ctx, &massagrpc.GetStatusRequest{})
	if err != nil {
		return nil, err
	}
	return res.Status, nil
}
