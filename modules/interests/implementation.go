package interests

import (
	"context"
	"errors"

	"google.golang.org/grpc"

	"github.com/FACorreiaa/loci-proto/core"
	c "github.com/FACorreiaa/loci-proto/modules/interests/generated"
	"github.com/FACorreiaa/loci-proto/utils"
)

type Broker struct {
	serverAddr string
	conn       *grpc.ClientConn
	client     c.InterestsServiceClient
}

var (
	_ c.InterestsServiceClient = (*Broker)(nil)
	_ core.Broker              = (*Broker)(nil)
)

func NewBroker(serverAddr string) (*Broker, error) {
	b := new(Broker)
	b.serverAddr = serverAddr

	if b.serverAddr == "" {
		return nil, errors.New("null routed upstream host")
	}

	return b, nil
}

func (b *Broker) NewConnection() (*grpc.ClientConn, error) {
	conn, err := utils.NewConnection(b.serverAddr)
	if err != nil {
		return nil, errors.New("could not open connection")
	}

	b.conn = conn
	b.client = c.NewInterestsServiceClient(b.conn)

	return b.conn, nil
}

func (b *Broker) GetAddress() string {
	return b.serverAddr
}

func (b *Broker) GetAllInterests(ctx context.Context, in *c.GetAllInterestsRequest, opts ...grpc.CallOption) (*c.GetAllInterestsResponse, error) {
	return b.client.GetAllInterests(ctx, in, opts...)
}

func (b *Broker) CreateInterest(ctx context.Context, in *c.CreateInterestRequest, opts ...grpc.CallOption) (*c.CreateInterestResponse, error) {
	return b.client.CreateInterest(ctx, in, opts...)
}

func (b *Broker) UpdateInterest(ctx context.Context, in *c.UpdateInterestRequest, opts ...grpc.CallOption) (*c.UpdateInterestResponse, error) {
	return b.client.UpdateInterest(ctx, in, opts...)
}

func (b *Broker) RemoveInterest(ctx context.Context, in *c.RemoveInterestRequest, opts ...grpc.CallOption) (*c.RemoveInterestResponse, error) {
	return b.client.RemoveInterest(ctx, in, opts...)
}
