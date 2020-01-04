package dialer

import (
	"context"

	"google.golang.org/grpc/credentials"
)

// Credentials holds per-rpc metadata for the gRPC clients
type Credentials struct {
	key, value string
	secure     bool
}

// NewCredentials instantiates a new credentials container that uses an arbitrary key/value combination.
func NewCredentials(key, value string, withTransportSecurity bool) credentials.PerRPCCredentials {
	return Credentials{key, value, withTransportSecurity}
}

// GetRequestMetadata retrieves relevant metadata
func (c Credentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		c.key: c.value,
	}, nil
}

// RequireTransportSecurity indicates that transport security is required
func (c Credentials) RequireTransportSecurity() bool { return c.secure }
