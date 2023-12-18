package grpc

import (
	"context"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/RaphaelL2e/golang-toolkit/constant"
)

type Server struct {
	*grpc.Server
	timeout            time.Duration
	baseCtx            context.Context
	address            string
	unaryInterceptors  []grpc.UnaryServerInterceptor
	streamInterceptors []grpc.StreamServerInterceptor
	grpcOptions        []grpc.ServerOption
}

// NewServer create a gRPC server by options
func NewServer(opts ...ServerOption) *Server {
	srv := &Server{
		baseCtx: context.Background(),
		address: ":0",
		timeout: constant.GrpcTimeout,
	}
	for _, o := range opts {
		o(srv)
	}
	var grpcOptions []grpc.ServerOption
	if len(srv.unaryInterceptors) > 0 {
		grpcOptions = append(grpcOptions, grpc.ChainUnaryInterceptor(srv.unaryInterceptors...))
	}
	if len(srv.streamInterceptors) > 0 {
		grpcOptions = append(grpcOptions, grpc.ChainStreamInterceptor(srv.streamInterceptors...))
	}
	grpcOptions = append(grpcOptions, srv.grpcOptions...)
	srv.Server = grpc.NewServer(grpcOptions...)
	reflection.Register(srv.Server)
	return srv
}

// Start the gRPC server
func (s *Server) Start(ctx context.Context) error {
	lis, err := net.Listen(constant.GrpcNetwork, s.address)
	if err != nil {
		return err
	}
	s.baseCtx = ctx
	return s.Serve(lis)
}

// Stop the gRPC server
func (s *Server) Stop(ctx context.Context) error {
	s.GracefulStop()
	return nil
}
