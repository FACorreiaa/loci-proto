package user

import (
	"context"
	"errors"

	"google.golang.org/grpc"

	"github.com/FACorreiaa/loci-proto/core"
	c "github.com/FACorreiaa/loci-proto/modules/user/generated"
	"github.com/FACorreiaa/loci-proto/utils"
)

type Broker struct {
	serverAddr string
	conn       *grpc.ClientConn
	client     c.UserServiceClient
}

var (
	_ c.UserServiceClient = (*Broker)(nil)
	_ core.Broker         = (*Broker)(nil)
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
	b.client = c.NewUserServiceClient(b.conn)

	return b.conn, nil
}

func (b *Broker) GetAddress() string {
	return b.serverAddr
}

// UserService method implementations
func (b *Broker) GetUserProfile(ctx context.Context, in *c.GetUserProfileRequest, opts ...grpc.CallOption) (*c.GetUserProfileResponse, error) {
	return b.client.GetUserProfile(ctx, in, opts...)
}

func (b *Broker) UpdateUserProfile(ctx context.Context, in *c.UpdateUserProfileRequest, opts ...grpc.CallOption) (*c.UpdateUserProfileResponse, error) {
	return b.client.UpdateUserProfile(ctx, in, opts...)
}

// Search profile management
func (b *Broker) GetSearchProfiles(ctx context.Context, in *c.GetSearchProfilesRequest, opts ...grpc.CallOption) (*c.GetSearchProfilesResponse, error) {
	return b.client.GetSearchProfiles(ctx, in, opts...)
}

func (b *Broker) GetSearchProfile(ctx context.Context, in *c.GetSearchProfileRequest, opts ...grpc.CallOption) (*c.GetSearchProfileResponse, error) {
	return b.client.GetSearchProfile(ctx, in, opts...)
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

func (b *Broker) GetDefaultProfile(ctx context.Context, in *c.GetDefaultProfileRequest, opts ...grpc.CallOption) (*c.GetDefaultProfileResponse, error) {
	return b.client.GetDefaultProfile(ctx, in, opts...)
}

func (b *Broker) SetDefaultProfile(ctx context.Context, in *c.SetDefaultProfileRequest, opts ...grpc.CallOption) (*c.SetDefaultProfileResponse, error) {
	return b.client.SetDefaultProfile(ctx, in, opts...)
}

// Interests management
func (b *Broker) GetInterests(ctx context.Context, in *c.GetInterestsRequest, opts ...grpc.CallOption) (*c.GetInterestsResponse, error) {
	return b.client.GetInterests(ctx, in, opts...)
}

func (b *Broker) CreateInterest(ctx context.Context, in *c.CreateInterestRequest, opts ...grpc.CallOption) (*c.CreateInterestResponse, error) {
	return b.client.CreateInterest(ctx, in, opts...)
}

func (b *Broker) UpdateInterest(ctx context.Context, in *c.UpdateInterestRequest, opts ...grpc.CallOption) (*c.UpdateInterestResponse, error) {
	return b.client.UpdateInterest(ctx, in, opts...)
}

func (b *Broker) DeleteInterest(ctx context.Context, in *c.DeleteInterestRequest, opts ...grpc.CallOption) (*c.DeleteInterestResponse, error) {
	return b.client.DeleteInterest(ctx, in, opts...)
}

// Tags management
func (b *Broker) GetTags(ctx context.Context, in *c.GetTagsRequest, opts ...grpc.CallOption) (*c.GetTagsResponse, error) {
	return b.client.GetTags(ctx, in, opts...)
}

func (b *Broker) GetTag(ctx context.Context, in *c.GetTagRequest, opts ...grpc.CallOption) (*c.GetTagResponse, error) {
	return b.client.GetTag(ctx, in, opts...)
}

func (b *Broker) CreateTag(ctx context.Context, in *c.CreateTagRequest, opts ...grpc.CallOption) (*c.CreateTagResponse, error) {
	return b.client.CreateTag(ctx, in, opts...)
}

func (b *Broker) UpdateTag(ctx context.Context, in *c.UpdateTagRequest, opts ...grpc.CallOption) (*c.UpdateTagResponse, error) {
	return b.client.UpdateTag(ctx, in, opts...)
}

func (b *Broker) DeleteTag(ctx context.Context, in *c.DeleteTagRequest, opts ...grpc.CallOption) (*c.DeleteTagResponse, error) {
	return b.client.DeleteTag(ctx, in, opts...)
}