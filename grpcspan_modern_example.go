package grpcspan

import (
	"google.golang.org/grpc/stats"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	"github.com/FACorreiaa/go-poi-au-suggestions/protocol/grpc/middleware"
)

// Handlers wraps OpenTelemetry gRPC stats handlers (recommended).
// We use a wrapper for this as the OpenTelemetry ecosystem changes
// *very* frequently, so we want to contain this change to a single point
// rather than having to update the servers and clients individually.
func Handlers() (middleware.ClientHandler, middleware.ServerHandler) {
	clientHandler := middleware.ClientHandler{
		Handler: otelgrpc.NewClientHandler(),
	}

	serverHandler := middleware.ServerHandler{
		Handler: otelgrpc.NewServerHandler(),
	}

	return clientHandler, serverHandler
}

// Deprecated: Use Handlers() instead
// Interceptors wraps OpenTelemetry gRPC interceptors (deprecated).
// The interceptor functions no longer exist in otelgrpc v0.62.0+
func Interceptors() (middleware.ClientInterceptor, middleware.ServerInterceptor) {
	// Return empty interceptors since the modern approach uses stats handlers
	return middleware.ClientInterceptor{}, middleware.ServerInterceptor{}
}

// NewClientHandler creates OpenTelemetry client stats handler with options
func NewClientHandler(opts ...otelgrpc.Option) stats.Handler {
	return otelgrpc.NewClientHandler(opts...)
}

// NewServerHandler creates OpenTelemetry server stats handler with options
func NewServerHandler(opts ...otelgrpc.Option) stats.Handler {
	return otelgrpc.NewServerHandler(opts...)
}

// Example usage functions for migration

// CreateClientConnection demonstrates how to create a gRPC client with OpenTelemetry
func CreateClientConnection(target string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	clientHandler, _ := Handlers()
	
	// Add OpenTelemetry stats handler to the dial options
	opts = append(opts,
		grpc.WithStatsHandler(clientHandler.Handler),
	)
	
	return grpc.NewClient(target, opts...)
}

// CreateServer demonstrates how to create a gRPC server with OpenTelemetry
func CreateServer(opts ...grpc.ServerOption) *grpc.Server {
	_, serverHandler := Handlers()
	
	// Add OpenTelemetry stats handler to the server options
	opts = append(opts,
		grpc.StatsHandler(serverHandler.Handler),
	)
	
	return grpc.NewServer(opts...)
}