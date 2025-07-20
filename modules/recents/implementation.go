package recents

import (
	"context"
	"errors"

	"google.golang.org/grpc"

	"github.com/FACorreiaa/loci-proto/core"
	c "github.com/FACorreiaa/loci-proto/modules/recents/generated"
	"github.com/FACorreiaa/loci-proto/utils"
)

type Broker struct {
	serverAddr string
	conn       *grpc.ClientConn
	client     c.RecentsServiceClient
}

var (
	_ c.RecentsServiceClient = (*Broker)(nil)
	_ core.Broker            = (*Broker)(nil)
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
	b.client = c.NewRecentsServiceClient(b.conn)

	return b.conn, nil
}

func (b *Broker) GetAddress() string {
	return b.serverAddr
}

// RecentsService method implementations
func (b *Broker) GetRecentInteractions(ctx context.Context, in *c.GetRecentInteractionsRequest, opts ...grpc.CallOption) (*c.GetRecentInteractionsResponse, error) {
	return b.client.GetRecentInteractions(ctx, in, opts...)
}

func (b *Broker) GetCityInteractions(ctx context.Context, in *c.GetCityInteractionsRequest, opts ...grpc.CallOption) (*c.GetCityInteractionsResponse, error) {
	return b.client.GetCityInteractions(ctx, in, opts...)
}

func (b *Broker) RecordInteraction(ctx context.Context, in *c.RecordInteractionRequest, opts ...grpc.CallOption) (*c.RecordInteractionResponse, error) {
	return b.client.RecordInteraction(ctx, in, opts...)
}

func (b *Broker) GetInteractionHistory(ctx context.Context, in *c.GetInteractionHistoryRequest, opts ...grpc.CallOption) (*c.GetInteractionHistoryResponse, error) {
	return b.client.GetInteractionHistory(ctx, in, opts...)
}

func (b *Broker) GetFrequentPlaces(ctx context.Context, in *c.GetFrequentPlacesRequest, opts ...grpc.CallOption) (*c.GetFrequentPlacesResponse, error) {
	return b.client.GetFrequentPlaces(ctx, in, opts...)
}