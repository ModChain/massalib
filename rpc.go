package massalib

import (
	"context"
	"fmt"
	"io"
	"log"

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

func (rpc *RPC) GetSlotTransfers(ctx context.Context, finality massagrpc.FinalityLevel) (chan *massagrpc.NewSlotTransfersResponse, error) {
	bidi, err := rpc.pub.NewSlotTransfers(ctx)
	if err != nil {
		return nil, fmt.Errorf("new slottransfers failed: %w", err)
	}
	//if err := bidi.CloseSend(); err != nil {
	//		return nil, fmt.Errorf("closeSend failed: %w", err)
	//	}

	ch := make(chan *massagrpc.NewSlotTransfersResponse)
	go func() {
		defer close(ch)
		for {
			resp, err := bidi.Recv()
			if err == io.EOF {
				log.Printf("got EOF")
				return
			}
			if err != nil {
				log.Printf("MASSA error on recv: %s", err)
				return
			}
			ch <- resp
		}
	}()

	return ch, nil
}
