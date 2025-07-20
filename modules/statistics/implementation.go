package statistics

import (
	"context"
	"errors"

	"google.golang.org/grpc"

	"github.com/FACorreiaa/loci-proto/core"
	c "github.com/FACorreiaa/loci-proto/modules/statistics/generated"
	"github.com/FACorreiaa/loci-proto/utils"
)

type Broker struct {
	serverAddr string
	conn       *grpc.ClientConn
	client     c.StatisticsServiceClient
}

var (
	_ c.StatisticsServiceClient = (*Broker)(nil)
	_ core.Broker               = (*Broker)(nil)
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
	b.client = c.NewStatisticsServiceClient(b.conn)

	return b.conn, nil
}

func (b *Broker) GetAddress() string {
	return b.serverAddr
}

// StatisticsService method implementations
func (b *Broker) GetMainPageStatistics(ctx context.Context, in *c.GetMainPageStatisticsRequest, opts ...grpc.CallOption) (*c.GetMainPageStatisticsResponse, error) {
	return b.client.GetMainPageStatistics(ctx, in, opts...)
}

// Streaming method
func (b *Broker) StreamMainPageStatistics(ctx context.Context, in *c.StreamMainPageStatisticsRequest, opts ...grpc.CallOption) (c.StatisticsService_StreamMainPageStatisticsClient, error) {
	return b.client.StreamMainPageStatistics(ctx, in, opts...)
}

func (b *Broker) GetDetailedPOIStatistics(ctx context.Context, in *c.GetDetailedPOIStatisticsRequest, opts ...grpc.CallOption) (*c.GetDetailedPOIStatisticsResponse, error) {
	return b.client.GetDetailedPOIStatistics(ctx, in, opts...)
}

func (b *Broker) GetLandingPageStatistics(ctx context.Context, in *c.GetLandingPageStatisticsRequest, opts ...grpc.CallOption) (*c.GetLandingPageStatisticsResponse, error) {
	return b.client.GetLandingPageStatistics(ctx, in, opts...)
}

func (b *Broker) GetUserActivityAnalytics(ctx context.Context, in *c.GetUserActivityAnalyticsRequest, opts ...grpc.CallOption) (*c.GetUserActivityAnalyticsResponse, error) {
	return b.client.GetUserActivityAnalytics(ctx, in, opts...)
}

func (b *Broker) GetSystemAnalytics(ctx context.Context, in *c.GetSystemAnalyticsRequest, opts ...grpc.CallOption) (*c.GetSystemAnalyticsResponse, error) {
	return b.client.GetSystemAnalytics(ctx, in, opts...)
}