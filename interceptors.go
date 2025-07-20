package proto

import (
	"google.golang.org/grpc"

	"github.com/FACorreiaa/loci-proto/middleware"
)

// Handlers wraps OpenTelemetry gRPC stats handlers (recommended).
// We use a wrapper for this as the OpenTelemetry ecosystem changes
// *very* frequently, so we want to contain this change to a single point
// rather than having to update the servers and clients individually.
func Handlers() (middleware.ClientHandler, middleware.ServerHandler) {
	return middleware.OpenTelemetryHandlers()
}

// Deprecated: Use Handlers() instead
// Interceptors wraps OpenTelemetry gRPC interceptors (deprecated).
func Interceptors() (middleware.ClientInterceptor, middleware.ServerInterceptor) {
	return middleware.OpenTelemetryInterceptors()
}

// NewClientConn creates a new gRPC client connection with OpenTelemetry stats handler
func NewClientConn(target string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	clientHandler, _ := Handlers()
	
	// Add OpenTelemetry stats handler to the dial options
	opts = append(opts,
		grpc.WithStatsHandler(clientHandler.Handler),
	)
	
	return grpc.NewClient(target, opts...)
}

// NewServer creates a new gRPC server with OpenTelemetry stats handler
func NewServer(opts ...grpc.ServerOption) *grpc.Server {
	_, serverHandler := Handlers()
	
	// Add OpenTelemetry stats handler to the server options
	opts = append(opts,
		grpc.StatsHandler(serverHandler.Handler),
	)
	
	return grpc.NewServer(opts...)
}

// Deprecated: Use NewClientConn() instead
// NewClientConnWithInterceptors creates a new gRPC client connection with OpenTelemetry interceptors (deprecated)
func NewClientConnWithInterceptors(target string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	clientInterceptor, _ := Interceptors()
	
	// Add empty interceptors for backwards compatibility
	opts = append(opts,
		grpc.WithUnaryInterceptor(clientInterceptor.Unary),
		grpc.WithStreamInterceptor(clientInterceptor.Stream),
	)
	
	return grpc.Dial(target, opts...)
}

// Deprecated: Use NewServer() instead
// NewServerWithInterceptors creates a new gRPC server with OpenTelemetry interceptors (deprecated)
func NewServerWithInterceptors(opts ...grpc.ServerOption) *grpc.Server {
	_, serverInterceptor := Interceptors()
	
	// Add empty interceptors for backwards compatibility
	opts = append(opts,
		grpc.UnaryInterceptor(serverInterceptor.Unary),
		grpc.StreamInterceptor(serverInterceptor.Stream),
	)
	
	return grpc.NewServer(opts...)
}