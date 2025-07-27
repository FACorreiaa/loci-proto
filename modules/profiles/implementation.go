package profiles

import (
	"context"
	"errors"

	"google.golang.org/grpc"

	"github.com/FACorreiaa/loci-proto/core"
	c "github.com/FACorreiaa/loci-proto/modules/profiles/generated"
	"github.com/FACorreiaa/loci-proto/utils"
)

type Broker struct {
	serverAddr string
	conn       *grpc.ClientConn
	client     c.ProfilesServiceClient
}

var (
	_ c.ProfilesServiceClient = (*Broker)(nil)
	_ core.Broker             = (*Broker)(nil)
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
	b.client = c.NewProfilesServiceClient(b.conn)

	return b.conn, nil
}

func (b *Broker) GetAddress() string {
	return b.serverAddr
}

func (b *Broker) GetSearchProfiles(ctx context.Context, in *c.GetSearchProfilesRequest, opts ...grpc.CallOption) (*c.GetSearchProfilesResponse, error) {
	return b.client.GetSearchProfiles(ctx, in, opts...)
}

func (b *Broker) GetSearchProfile(ctx context.Context, in *c.GetSearchProfileRequest, opts ...grpc.CallOption) (*c.GetSearchProfileResponse, error) {
	return b.client.GetSearchProfile(ctx, in, opts...)
}

func (b *Broker) GetDefaultSearchProfile(ctx context.Context, in *c.GetDefaultSearchProfileRequest, opts ...grpc.CallOption) (*c.GetDefaultSearchProfileResponse, error) {
	return b.client.GetDefaultSearchProfile(ctx, in, opts...)
}

func (b *Broker) CreateSearchProfile(ctx context.Context, in *c.CreateSearchProfileRequest, opts ...grpc.CallOption) (*c.CreateSearchProfileResponse, error) {
	return b.client.CreateSearchProfile(ctx, in, opts...)
}

func (b *Broker) UpdateSearchProfile(ctx context.Context, in *c.UpdateSearchProfileRequest, opts ...grpc.CallOption) (*c.UpdateSearchProfileResponse, error) {
	return b.client.UpdateSearchProfile(ctx, in, opts...)
}

func (b *Broker) DeleteSearchProfile(ctx context.Context, in *c.DeleteSearchProfileRequest, opts ...grpc.CallOption) (*c.DeleteSearchProfileResponse, error) {
	return b.client.DeleteSearchProfile(ctx, in, opts...)
}

func (b *Broker) SetDefaultSearchProfile(ctx context.Context, in *c.SetDefaultSearchProfileRequest, opts ...grpc.CallOption) (*c.SetDefaultSearchProfileResponse, error) {
	return b.client.SetDefaultSearchProfile(ctx, in, opts...)
}
