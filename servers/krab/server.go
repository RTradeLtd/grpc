// Package krab provides a microservice to handle ipfs key management
package krab

import (
	"context"
	"fmt"
	"net"

	"github.com/RTradeLtd/config"
	pb "github.com/RTradeLtd/grpc/krab"
	"github.com/RTradeLtd/grpc/middleware"
	"github.com/RTradeLtd/rtfs/krab"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	ci "github.com/libp2p/go-libp2p-crypto"
)

// Server is the backend for Krab
type Server struct {
	krab *krab.Krab
	pb.ServiceServer
}

// NewServer is used to create, and run a krab keystore server
func NewServer(listenAddr, protocol string, cfg *config.TemporalConfig) error {
	lis, err := net.Listen(protocol, listenAddr)
	if err != nil {
		return err
	}
	// setup authentication interceptor
	unaryIntercept, streamInterceptor := middleware.NewServerInterceptors(cfg.Endpoints.Krab.AuthKey)
	// setup server options
	serverOpts := []grpc.ServerOption{
		grpc_middleware.WithUnaryServerChain(
			unaryIntercept,
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor))),
		grpc_middleware.WithStreamServerChain(
			streamInterceptor,
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor))),
	}
	// setup tls configuration if available
	if cfg.Endpoints.Krab.TLS.CertPath != "" {
		creds, err := credentials.NewServerTLSFromFile(
			cfg.Krab.TLS.CertPath,
			cfg.Krab.TLS.KeyFile,
		)
		if err != nil {
			return err
		}
		serverOpts = append(serverOpts, grpc.Creds(creds))
	}
	// create grpc server
	gServer := grpc.NewServer(serverOpts...)
	// setup krab backend
	kb, err := krab.NewKrab(krab.Opts{Passphrase: cfg.Endpoints.Krab.KeystorePassword, DSPath: cfg.IPFS.KeystorePath, ReadOnly: false})
	if err != nil {
		return err
	}
	defer func() {
		if err := kb.Close(); err != nil {
			fmt.Println("failed to properly close datastore connection ", err.Error())
		}
	}()
	server := &Server{
		krab: kb,
	}
	defer gServer.GracefulStop()
	pb.RegisterServiceServer(gServer, server)
	if err = gServer.Serve(lis); err != nil {
		return err
	}
	// register gRPC services
	return nil
}

// KeyGet is used to retrieve a private key by searching for its name
func (s *Server) KeyGet(ctx context.Context, req *pb.KeyGet) (*pb.Response, error) {
	pk, err := s.krab.Get(req.Name)
	if err != nil {
		return nil, err
	}
	keyBytes, err := pk.Bytes()
	if err != nil {
		return nil, err
	}
	return &pb.Response{
		Status:     "get succesful",
		PrivateKey: keyBytes,
	}, nil
}

// KeyPut is used to store a new private key
func (s *Server) KeyPut(ctx context.Context, req *pb.KeyPut) (*pb.Response, error) {
	pk, err := ci.UnmarshalPrivateKey(req.PrivateKey)
	if err != nil {
		return nil, err
	}
	if err := s.krab.Put(req.Name, pk); err != nil {
		return nil, err
	}
	return &pb.Response{
		Status: "key store",
	}, nil
}