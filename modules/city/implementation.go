package city

import (
	"context"
	"errors"

	"google.golang.org/grpc"

	"github.com/FACorreiaa/loci-proto/core"
	c "github.com/FACorreiaa/loci-proto/modules/city/generated"
	"github.com/FACorreiaa/loci-proto/utils"
)

type Broker struct {
	serverAddr string
	conn       *grpc.ClientConn
	client     c.CityServiceClient
}

var (
	_ c.CityServiceClient = (*Broker)(nil)
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
	b.client = c.NewCityServiceClient(b.conn)

	return b.conn, nil
}

func (b *Broker) GetAddress() string {
	return b.serverAddr
}

// CityService method implementations
func (b *Broker) GetCities(ctx context.Context, in *c.GetCitiesRequest, opts ...grpc.CallOption) (*c.GetCitiesResponse, error) {
	return b.client.GetCities(ctx, in, opts...)
}

func (b *Broker) GetCity(ctx context.Context, in *c.GetCityRequest, opts ...grpc.CallOption) (*c.GetCityResponse, error) {
	return b.client.GetCity(ctx, in, opts...)
}

func (b *Broker) SearchCities(ctx context.Context, in *c.SearchCitiesRequest, opts ...grpc.CallOption) (*c.SearchCitiesResponse, error) {
	return b.client.SearchCities(ctx, in, opts...)
}

func (b *Broker) GetCityStatistics(ctx context.Context, in *c.GetCityStatisticsRequest, opts ...grpc.CallOption) (*c.GetCityStatisticsResponse, error) {
	return b.client.GetCityStatistics(ctx, in, opts...)
}