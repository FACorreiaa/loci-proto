package ai_poi_service

import (
	"context"
	"errors"

	"google.golang.org/grpc"

	"github.com/FACorreiaa/loci-proto/core"
	c "github.com/FACorreiaa/loci-proto/modules/ai_poi_service/generated"
	"github.com/FACorreiaa/loci-proto/utils"
)

type Broker struct {
	serverAddr string
	conn       *grpc.ClientConn
	client     c.AiPoiServiceClient
}

var (
	_ c.AiPoiServiceClient = (*Broker)(nil)
	_ core.Broker          = (*Broker)(nil)
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
	b.client = c.NewAiPoiServiceClient(b.conn)

	return b.conn, nil
}

func (b *Broker) GetAddress() string {
	return b.serverAddr
}

// AiPoiService method implementations
func (b *Broker) HealthCheck(ctx context.Context, in *c.HealthCheckRequest, opts ...grpc.CallOption) (*c.HealthCheckResponse, error) {
	return b.client.HealthCheck(ctx, in, opts...)
}

func (b *Broker) GetServiceInfo(ctx context.Context, in *c.GetServiceInfoRequest, opts ...grpc.CallOption) (*c.GetServiceInfoResponse, error) {
	return b.client.GetServiceInfo(ctx, in, opts...)
}

func (b *Broker) GetFeatureFlags(ctx context.Context, in *c.GetFeatureFlagsRequest, opts ...grpc.CallOption) (*c.GetFeatureFlagsResponse, error) {
	return b.client.GetFeatureFlags(ctx, in, opts...)
}