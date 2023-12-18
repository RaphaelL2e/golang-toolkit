package grpc

import (
	"time"

	"google.golang.org/grpc"
)

// ServerOption is gRPC server option.
type ServerOption func(s *Server)

// Address with server address
func Address(address string) ServerOption {
	return func(s *Server) {
		s.address = address
	}
}

// Timeout with server timeout
func Timeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.timeout = timeout
	}
}

// UnaryInterceptors with unary interceptors
func UnaryInterceptors(unaryInterceptors ...grpc.UnaryServerInterceptor) ServerOption {
	return func(s *Server) {
		s.unaryInterceptors = append(s.unaryInterceptors, unaryInterceptors...)
	}
}

func StreamInterceptors(streamInterceptors ...grpc.StreamServerInterceptor) ServerOption {
	return func(s *Server) {
		s.streamInterceptors = append(s.streamInterceptors, streamInterceptors...)
	}
}

func Options(options ...grpc.ServerOption) ServerOption {
	return func(s *Server) {
		s.grpcOptions = options
	}
}
