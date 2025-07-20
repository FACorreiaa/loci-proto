package customer

import (
	"context"
	"errors"

	"google.golang.org/grpc"

	"github.com/FACorreiaa/loci-proto/core"
	c "github.com/FACorreiaa/loci-proto/modules/customer/generated"
	"github.com/FACorreiaa/loci-proto/utils"
)

type Broker struct {
	serverAddr string
	conn       *grpc.ClientConn
	client     c.CustomerClient
}

var (
	_ c.CustomerClient = (*Broker)(nil)
	_ core.Broker      = (*Broker)(nil)
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
	b.client = c.NewCustomerClient(b.conn)

	return b.conn, nil
}

func (b *Broker) GetAddress() string {
	return b.serverAddr
}

func (b *Broker) GetCustomer(ctx context.Context, in *c.GetCustomerReq, opts ...grpc.CallOption) (*c.GetCustomerRes, error) {
	return b.client.GetCustomer(ctx, in, opts...)
}

func (b *Broker) CreateCustomer(ctx context.Context, in *c.CreateCustomerReq, opts ...grpc.CallOption) (*c.CreateCustomerRes, error) {
	return b.client.CreateCustomer(ctx, in, opts...)
}

func (b *Broker) UpdateCustomer(ctx context.Context, in *c.UpdateCustomerReq, opts ...grpc.CallOption) (*c.UpdateCustomerRes, error) {
	return b.client.UpdateCustomer(ctx, in, opts...)
}

func (b *Broker) DeleteCustomer(ctx context.Context, in *c.DeleteCustomerReq, opts ...grpc.CallOption) (*c.NilRes, error) {
	return b.client.DeleteCustomer(ctx, in, opts...)
}
