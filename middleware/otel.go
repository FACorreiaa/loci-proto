package middleware

import (
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
)

// OpenTelemetryHandlers creates OpenTelemetry gRPC stats handlers.
// We use a wrapper for this as the OpenTelemetry ecosystem changes
// *very* frequently, so we want to contain this change to a single point
// rather than having to update the servers and clients individually.
func OpenTelemetryHandlers() (ClientHandler, ServerHandler) {
	clientHandler := ClientHandler{
		Handler: otelgrpc.NewClientHandler(),
	}

	serverHandler := ServerHandler{
		Handler: otelgrpc.NewServerHandler(),
	}

	return clientHandler, serverHandler
}

// Deprecated: Use OpenTelemetryHandlers() instead
// OpenTelemetryInterceptors is kept for backwards compatibility
func OpenTelemetryInterceptors() (ClientInterceptor, ServerInterceptor) {
	// Return empty interceptors since the modern approach uses stats handlers
	return ClientInterceptor{}, ServerInterceptor{}
}

// NewOtelClientHandler creates OpenTelemetry client stats handler
func NewOtelClientHandler() ClientHandler {
	client, _ := OpenTelemetryHandlers()
	return client
}

// NewOtelServerHandler creates OpenTelemetry server stats handler
func NewOtelServerHandler() ServerHandler {
	_, server := OpenTelemetryHandlers()
	return server
}