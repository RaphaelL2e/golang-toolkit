package http

import (
	"context"
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/RaphaelL2e/golang-toolkit/constant"
)

type Server struct {
	*http.Server
	address string
	router  *mux.Router
}

func NewServer(opts ...ServerOption) *Server {
	srv := &Server{
		address: ":0",
	}
	for _, o := range opts {
		o(srv)
	}
	srv.router = mux.NewRouter()

	srv.Server = &http.Server{
		Addr:              srv.address,
		Handler:           srv.Handler,
		ReadHeaderTimeout: constant.HttpTimeout,
	}
	return srv
}

func (s *Server) Start(ctx context.Context) error {
	err := s.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return s.Shutdown(ctx)
}

// HandlePrefix registers a handler for a given prefix.
//
// This is equivalent to:
//
//	router.Handle(prefix, h)
//	router.Handle(prefix+"/", h)
func (s *Server) HandlePrefix(prefix string, h http.Handler) {
	s.router.PathPrefix(prefix).Handler(h)
}
