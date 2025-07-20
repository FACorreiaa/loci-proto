package middleware

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/stats"
)

// ClientInterceptor holds gRPC client interceptors (deprecated)
type ClientInterceptor struct {
	Unary  grpc.UnaryClientInterceptor
	Stream grpc.StreamClientInterceptor
}

// ServerInterceptor holds gRPC server interceptors (deprecated)
type ServerInterceptor struct {
	Unary  grpc.UnaryServerInterceptor
	Stream grpc.StreamServerInterceptor
}

// ClientHandler holds gRPC client stats handler (recommended)
type ClientHandler struct {
	Handler stats.Handler
}

// ServerHandler holds gRPC server stats handler (recommended)
type ServerHandler struct {
	Handler stats.Handler
}