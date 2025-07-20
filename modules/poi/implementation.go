package poi

import (
	"context"
	"errors"

	"google.golang.org/grpc"

	"github.com/FACorreiaa/loci-proto/core"
	c "github.com/FACorreiaa/loci-proto/modules/poi/generated"
	"github.com/FACorreiaa/loci-proto/utils"
)

type Broker struct {
	serverAddr string
	conn       *grpc.ClientConn
	client     c.POIServiceClient
}

var (
	_ c.POIServiceClient = (*Broker)(nil)
	_ core.Broker        = (*Broker)(nil)
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
	b.client = c.NewPOIServiceClient(b.conn)

	return b.conn, nil
}

func (b *Broker) GetAddress() string {
	return b.serverAddr
}

// POIService method implementations
func (b *Broker) GetPOIsByCity(ctx context.Context, in *c.GetPOIsByCityRequest, opts ...grpc.CallOption) (*c.GetPOIsByCityResponse, error) {
	return b.client.GetPOIsByCity(ctx, in, opts...)
}

func (b *Broker) SearchPOIs(ctx context.Context, in *c.SearchPOIsRequest, opts ...grpc.CallOption) (*c.SearchPOIsResponse, error) {
	return b.client.SearchPOIs(ctx, in, opts...)
}

func (b *Broker) SearchPOIsSemantic(ctx context.Context, in *c.SearchPOIsSemanticRequest, opts ...grpc.CallOption) (*c.SearchPOIsSemanticResponse, error) {
	return b.client.SearchPOIsSemantic(ctx, in, opts...)
}

func (b *Broker) SearchPOIsSemanticByCity(ctx context.Context, in *c.SearchPOIsSemanticByCityRequest, opts ...grpc.CallOption) (*c.SearchPOIsSemanticResponse, error) {
	return b.client.SearchPOIsSemanticByCity(ctx, in, opts...)
}

func (b *Broker) SearchPOIsHybrid(ctx context.Context, in *c.SearchPOIsHybridRequest, opts ...grpc.CallOption) (*c.SearchPOIsHybridResponse, error) {
	return b.client.SearchPOIsHybrid(ctx, in, opts...)
}

func (b *Broker) GetNearbyRecommendations(ctx context.Context, in *c.GetNearbyRecommendationsRequest, opts ...grpc.CallOption) (*c.GetNearbyRecommendationsResponse, error) {
	return b.client.GetNearbyRecommendations(ctx, in, opts...)
}

func (b *Broker) DiscoverRestaurants(ctx context.Context, in *c.DiscoverRestaurantsRequest, opts ...grpc.CallOption) (*c.DiscoverRestaurantsResponse, error) {
	return b.client.DiscoverRestaurants(ctx, in, opts...)
}

func (b *Broker) DiscoverActivities(ctx context.Context, in *c.DiscoverActivitiesRequest, opts ...grpc.CallOption) (*c.DiscoverActivitiesResponse, error) {
	return b.client.DiscoverActivities(ctx, in, opts...)
}

func (b *Broker) DiscoverHotels(ctx context.Context, in *c.DiscoverHotelsRequest, opts ...grpc.CallOption) (*c.DiscoverHotelsResponse, error) {
	return b.client.DiscoverHotels(ctx, in, opts...)
}

func (b *Broker) DiscoverAttractions(ctx context.Context, in *c.DiscoverAttractionsRequest, opts ...grpc.CallOption) (*c.DiscoverAttractionsResponse, error) {
	return b.client.DiscoverAttractions(ctx, in, opts...)
}

func (b *Broker) AddToFavorites(ctx context.Context, in *c.AddToFavoritesRequest, opts ...grpc.CallOption) (*c.AddToFavoritesResponse, error) {
	return b.client.AddToFavorites(ctx, in, opts...)
}

func (b *Broker) RemoveFromFavorites(ctx context.Context, in *c.RemoveFromFavoritesRequest, opts ...grpc.CallOption) (*c.RemoveFromFavoritesResponse, error) {
	return b.client.RemoveFromFavorites(ctx, in, opts...)
}

func (b *Broker) GetFavorites(ctx context.Context, in *c.GetFavoritesRequest, opts ...grpc.CallOption) (*c.GetFavoritesResponse, error) {
	return b.client.GetFavorites(ctx, in, opts...)
}

func (b *Broker) GetItineraries(ctx context.Context, in *c.GetItinerariesRequest, opts ...grpc.CallOption) (*c.GetItinerariesResponse, error) {
	return b.client.GetItineraries(ctx, in, opts...)
}

func (b *Broker) GetItinerary(ctx context.Context, in *c.GetItineraryRequest, opts ...grpc.CallOption) (*c.GetItineraryResponse, error) {
	return b.client.GetItinerary(ctx, in, opts...)
}

func (b *Broker) UpdateItinerary(ctx context.Context, in *c.UpdateItineraryRequest, opts ...grpc.CallOption) (*c.UpdateItineraryResponse, error) {
	return b.client.UpdateItinerary(ctx, in, opts...)
}

func (b *Broker) GenerateEmbeddings(ctx context.Context, in *c.GenerateEmbeddingsRequest, opts ...grpc.CallOption) (*c.GenerateEmbeddingsResponse, error) {
	return b.client.GenerateEmbeddings(ctx, in, opts...)
}