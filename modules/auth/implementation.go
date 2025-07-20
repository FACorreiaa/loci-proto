package auth

import (
	"context"
	"errors"

	"google.golang.org/grpc"

	"github.com/FACorreiaa/loci-proto/core"
	c "github.com/FACorreiaa/loci-proto/modules/auth/generated"
	"github.com/FACorreiaa/loci-proto/utils"
)

type Broker struct {
	serverAddr string
	conn       *grpc.ClientConn
	client     c.AuthServiceClient
}

var (
	_ c.AuthServiceClient = (*Broker)(nil)
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
	b.client = c.NewAuthServiceClient(b.conn)

	return b.conn, nil
}

func (b *Broker) GetAddress() string {
	return b.serverAddr
}

// AuthService method implementations
func (b *Broker) Register(ctx context.Context, in *c.RegisterRequest, opts ...grpc.CallOption) (*c.RegisterResponse, error) {
	return b.client.Register(ctx, in, opts...)
}

func (b *Broker) Login(ctx context.Context, in *c.LoginRequest, opts ...grpc.CallOption) (*c.LoginResponse, error) {
	return b.client.Login(ctx, in, opts...)
}

func (b *Broker) RefreshToken(ctx context.Context, in *c.RefreshTokenRequest, opts ...grpc.CallOption) (*c.TokenResponse, error) {
	return b.client.RefreshToken(ctx, in, opts...)
}

func (b *Broker) Logout(ctx context.Context, in *c.LogoutRequest, opts ...grpc.CallOption) (*c.LogoutResponse, error) {
	return b.client.Logout(ctx, in, opts...)
}

func (b *Broker) ValidateSession(ctx context.Context, in *c.ValidateSessionRequest, opts ...grpc.CallOption) (*c.ValidateSessionResponse, error) {
	return b.client.ValidateSession(ctx, in, opts...)
}

func (b *Broker) UpdatePassword(ctx context.Context, in *c.UpdatePasswordRequest, opts ...grpc.CallOption) (*c.UpdatePasswordResponse, error) {
	return b.client.UpdatePassword(ctx, in, opts...)
}

func (b *Broker) GoogleLogin(ctx context.Context, in *c.GoogleLoginRequest, opts ...grpc.CallOption) (*c.GoogleLoginResponse, error) {
	return b.client.GoogleLogin(ctx, in, opts...)
}

func (b *Broker) GoogleCallback(ctx context.Context, in *c.GoogleCallbackRequest, opts ...grpc.CallOption) (*c.LoginResponse, error) {
	return b.client.GoogleCallback(ctx, in, opts...)
}