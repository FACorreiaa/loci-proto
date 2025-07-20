package chat

import (
	"context"
	"errors"

	"google.golang.org/grpc"

	"github.com/FACorreiaa/loci-proto/core"
	c "github.com/FACorreiaa/loci-proto/modules/chat/generated"
	"github.com/FACorreiaa/loci-proto/utils"
)

type Broker struct {
	serverAddr string
	conn       *grpc.ClientConn
	client     c.ChatServiceClient
}

var (
	_ c.ChatServiceClient = (*Broker)(nil)
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
	b.client = c.NewChatServiceClient(b.conn)

	return b.conn, nil
}

func (b *Broker) GetAddress() string {
	return b.serverAddr
}

// ChatService method implementations

// Streaming methods
func (b *Broker) StartChatStream(ctx context.Context, in *c.StartChatRequest, opts ...grpc.CallOption) (c.ChatService_StartChatStreamClient, error) {
	return b.client.StartChatStream(ctx, in, opts...)
}

func (b *Broker) ContinueChatStream(ctx context.Context, in *c.ContinueChatRequest, opts ...grpc.CallOption) (c.ChatService_ContinueChatStreamClient, error) {
	return b.client.ContinueChatStream(ctx, in, opts...)
}

func (b *Broker) FreeChatStream(ctx context.Context, in *c.FreeChatRequest, opts ...grpc.CallOption) (c.ChatService_FreeChatStreamClient, error) {
	return b.client.FreeChatStream(ctx, in, opts...)
}

// Non-streaming methods
func (b *Broker) GetChatSessions(ctx context.Context, in *c.GetChatSessionsRequest, opts ...grpc.CallOption) (*c.GetChatSessionsResponse, error) {
	return b.client.GetChatSessions(ctx, in, opts...)
}

func (b *Broker) SaveItinerary(ctx context.Context, in *c.SaveItineraryRequest, opts ...grpc.CallOption) (*c.SaveItineraryResponse, error) {
	return b.client.SaveItinerary(ctx, in, opts...)
}

func (b *Broker) GetSavedItineraries(ctx context.Context, in *c.GetSavedItinerariesRequest, opts ...grpc.CallOption) (*c.GetSavedItinerariesResponse, error) {
	return b.client.GetSavedItineraries(ctx, in, opts...)
}

func (b *Broker) RemoveItinerary(ctx context.Context, in *c.RemoveItineraryRequest, opts ...grpc.CallOption) (*c.RemoveItineraryResponse, error) {
	return b.client.RemoveItinerary(ctx, in, opts...)
}

func (b *Broker) GetPOIDetails(ctx context.Context, in *c.GetPOIDetailsRequest, opts ...grpc.CallOption) (*c.GetPOIDetailsResponse, error) {
	return b.client.GetPOIDetails(ctx, in, opts...)
}