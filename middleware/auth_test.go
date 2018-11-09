package middleware

import (
	"context"
	"testing"

	"google.golang.org/grpc/metadata"
)

func TestNewServerInterceptors(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		reqKey  string
		wantErr bool
	}{
		{"wrong key", args{"temporal"}, "not-temporal", true},
		{"correct key", args{"temporal"}, "temporal", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			unary, _ := NewServerInterceptors(tt.args.key)
			ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(
				AuthorizationKey, tt.reqKey,
			))
			if _, err := unary(ctx, nil, nil, nil); (err != nil) != tt.wantErr {
				t.Errorf("unary() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
