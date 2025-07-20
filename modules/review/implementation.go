package review

import (
	"context"
	"errors"

	"google.golang.org/grpc"

	"github.com/FACorreiaa/loci-proto/core"
	c "github.com/FACorreiaa/loci-proto/modules/review/generated"
	"github.com/FACorreiaa/loci-proto/utils"
)

type Broker struct {
	serverAddr string
	conn       *grpc.ClientConn
	client     c.ReviewServiceClient
}

var (
	_ c.ReviewServiceClient = (*Broker)(nil)
	_ core.Broker           = (*Broker)(nil)
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
	b.client = c.NewReviewServiceClient(b.conn)

	return b.conn, nil
}

func (b *Broker) GetAddress() string {
	return b.serverAddr
}

// ReviewService method implementations
func (b *Broker) CreateReview(ctx context.Context, in *c.CreateReviewRequest, opts ...grpc.CallOption) (*c.CreateReviewResponse, error) {
	return b.client.CreateReview(ctx, in, opts...)
}

func (b *Broker) GetPOIReviews(ctx context.Context, in *c.GetPOIReviewsRequest, opts ...grpc.CallOption) (*c.GetPOIReviewsResponse, error) {
	return b.client.GetPOIReviews(ctx, in, opts...)
}

func (b *Broker) GetReview(ctx context.Context, in *c.GetReviewRequest, opts ...grpc.CallOption) (*c.GetReviewResponse, error) {
	return b.client.GetReview(ctx, in, opts...)
}

func (b *Broker) UpdateReview(ctx context.Context, in *c.UpdateReviewRequest, opts ...grpc.CallOption) (*c.UpdateReviewResponse, error) {
	return b.client.UpdateReview(ctx, in, opts...)
}

func (b *Broker) DeleteReview(ctx context.Context, in *c.DeleteReviewRequest, opts ...grpc.CallOption) (*c.DeleteReviewResponse, error) {
	return b.client.DeleteReview(ctx, in, opts...)
}

func (b *Broker) GetUserReviews(ctx context.Context, in *c.GetUserReviewsRequest, opts ...grpc.CallOption) (*c.GetUserReviewsResponse, error) {
	return b.client.GetUserReviews(ctx, in, opts...)
}

func (b *Broker) LikeReview(ctx context.Context, in *c.LikeReviewRequest, opts ...grpc.CallOption) (*c.LikeReviewResponse, error) {
	return b.client.LikeReview(ctx, in, opts...)
}

func (b *Broker) ReportReview(ctx context.Context, in *c.ReportReviewRequest, opts ...grpc.CallOption) (*c.ReportReviewResponse, error) {
	return b.client.ReportReview(ctx, in, opts...)
}

func (b *Broker) GetReviewStatistics(ctx context.Context, in *c.GetReviewStatisticsRequest, opts ...grpc.CallOption) (*c.GetReviewStatisticsResponse, error) {
	return b.client.GetReviewStatistics(ctx, in, opts...)
}