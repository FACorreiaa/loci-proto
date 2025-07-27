package tags

import (
	"context"

	"github.com/FACorreiaa/loci-proto/core"
	"github.com/FACorreiaa/loci-proto/utils"

	"google.golang.org/grpc"

	"errors"

	c "github.com/FACorreiaa/loci-proto/modules/tags/generated"
)

type Broker struct {
	serverAddr string
	conn       *grpc.ClientConn
	client     c.TagsServiceClient
}

var (
	_ c.TagsServiceClient = (*Broker)(nil)
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
	b.client = c.NewTagsServiceClient(b.conn)

	return b.conn, nil
}

func (b *Broker) GetAddress() string {
	return b.serverAddr
}

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
