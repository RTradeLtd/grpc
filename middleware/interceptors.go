package middleware

import (
	"context"

	"google.golang.org/grpc"
)

// UnaryInterceptor is used to intercept unary requests
type UnaryInterceptor func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)

// StreamInterceptor is used to intercept stream requests
type StreamInterceptor func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error
