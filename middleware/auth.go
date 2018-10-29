package middleware

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

// NewServerInterceptors creates unary and stream interceptors that validate
// requests, for use with gRPC servers, using given key
func NewServerInterceptors(key string) (
	unaryInterceptor func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error),
	streamInterceptor func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error,
) {
	unaryInterceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		meta, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, grpc.Errorf(codes.Unauthenticated, "missing context metadata")
		}
		if len(meta["token"]) != 1 {
			return nil, grpc.Errorf(codes.Unauthenticated, "invalid token")
		}
		if meta["token"][0] != key {
			return nil, grpc.Errorf(codes.Unauthenticated, "invalid token")
		}

		return handler(ctx, req)
	}

	streamInterceptor = func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		meta, ok := metadata.FromIncomingContext(stream.Context())
		if !ok {
			return grpc.Errorf(codes.Unauthenticated, "missing context metadata")
		}
		if len(meta["token"]) != 1 {
			return grpc.Errorf(codes.Unauthenticated, "invalid token")
		}
		if meta["token"][0] != key {
			return grpc.Errorf(codes.Unauthenticated, "invalid token")
		}

		err := handler(srv, stream)
		return err
	}

	return
}
