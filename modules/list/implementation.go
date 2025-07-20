package list

import (
	"context"
	"errors"

	"google.golang.org/grpc"

	"github.com/FACorreiaa/loci-proto/core"
	c "github.com/FACorreiaa/loci-proto/modules/list/generated"
	"github.com/FACorreiaa/loci-proto/utils"
)

type Broker struct {
	serverAddr string
	conn       *grpc.ClientConn
	client     c.ListServiceClient
}

var (
	_ c.ListServiceClient = (*Broker)(nil)
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
	b.client = c.NewListServiceClient(b.conn)

	return b.conn, nil
}

func (b *Broker) GetAddress() string {
	return b.serverAddr
}

// ListService method implementations

// List management
func (b *Broker) CreateList(ctx context.Context, in *c.CreateListRequest, opts ...grpc.CallOption) (*c.CreateListResponse, error) {
	return b.client.CreateList(ctx, in, opts...)
}

func (b *Broker) GetLists(ctx context.Context, in *c.GetListsRequest, opts ...grpc.CallOption) (*c.GetListsResponse, error) {
	return b.client.GetLists(ctx, in, opts...)
}

func (b *Broker) GetList(ctx context.Context, in *c.GetListRequest, opts ...grpc.CallOption) (*c.GetListResponse, error) {
	return b.client.GetList(ctx, in, opts...)
}

func (b *Broker) UpdateList(ctx context.Context, in *c.UpdateListRequest, opts ...grpc.CallOption) (*c.UpdateListResponse, error) {
	return b.client.UpdateList(ctx, in, opts...)
}

func (b *Broker) DeleteList(ctx context.Context, in *c.DeleteListRequest, opts ...grpc.CallOption) (*c.DeleteListResponse, error) {
	return b.client.DeleteList(ctx, in, opts...)
}

// Itinerary management
func (b *Broker) CreateItinerary(ctx context.Context, in *c.CreateItineraryRequest, opts ...grpc.CallOption) (*c.CreateItineraryResponse, error) {
	return b.client.CreateItinerary(ctx, in, opts...)
}

// List item management
func (b *Broker) AddListItem(ctx context.Context, in *c.AddListItemRequest, opts ...grpc.CallOption) (*c.AddListItemResponse, error) {
	return b.client.AddListItem(ctx, in, opts...)
}

func (b *Broker) UpdateListItem(ctx context.Context, in *c.UpdateListItemRequest, opts ...grpc.CallOption) (*c.UpdateListItemResponse, error) {
	return b.client.UpdateListItem(ctx, in, opts...)
}

func (b *Broker) RemoveListItem(ctx context.Context, in *c.RemoveListItemRequest, opts ...grpc.CallOption) (*c.RemoveListItemResponse, error) {
	return b.client.RemoveListItem(ctx, in, opts...)
}

func (b *Broker) GetListItems(ctx context.Context, in *c.GetListItemsRequest, opts ...grpc.CallOption) (*c.GetListItemsResponse, error) {
	return b.client.GetListItems(ctx, in, opts...)
}

// Get items by content type
func (b *Broker) GetListRestaurants(ctx context.Context, in *c.GetListRestaurantsRequest, opts ...grpc.CallOption) (*c.GetListRestaurantsResponse, error) {
	return b.client.GetListRestaurants(ctx, in, opts...)
}

func (b *Broker) GetListHotels(ctx context.Context, in *c.GetListHotelsRequest, opts ...grpc.CallOption) (*c.GetListHotelsResponse, error) {
	return b.client.GetListHotels(ctx, in, opts...)
}

func (b *Broker) GetListItineraries(ctx context.Context, in *c.GetListItinerariesRequest, opts ...grpc.CallOption) (*c.GetListItinerariesResponse, error) {
	return b.client.GetListItineraries(ctx, in, opts...)
}

// Public list management
func (b *Broker) SavePublicList(ctx context.Context, in *c.SavePublicListRequest, opts ...grpc.CallOption) (*c.SavePublicListResponse, error) {
	return b.client.SavePublicList(ctx, in, opts...)
}

func (b *Broker) UnsaveList(ctx context.Context, in *c.UnsaveListRequest, opts ...grpc.CallOption) (*c.UnsaveListResponse, error) {
	return b.client.UnsaveList(ctx, in, opts...)
}

func (b *Broker) GetSavedLists(ctx context.Context, in *c.GetSavedListsRequest, opts ...grpc.CallOption) (*c.GetSavedListsResponse, error) {
	return b.client.GetSavedLists(ctx, in, opts...)
}

func (b *Broker) SearchPublicLists(ctx context.Context, in *c.SearchPublicListsRequest, opts ...grpc.CallOption) (*c.SearchPublicListsResponse, error) {
	return b.client.SearchPublicLists(ctx, in, opts...)
}