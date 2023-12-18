package daemon

import (
	"context"

	"go.uber.org/zap"
)

type Daemon interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

type Server struct {
	ctx    context.Context
	cancel context.CancelFunc
	d      Daemon
}

func NewServer(d Daemon) *Server {
	return &Server{d: d}
}

func (s *Server) Start(ctx context.Context) error {
	s.ctx, s.cancel = context.WithCancel(ctx)
	err := s.d.Start(s.ctx)
	if err != nil {
		return err
	}
	zap.L().Debug("daemon server started")
	<-s.ctx.Done()
	zap.L().Debug("daemon server ended")

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	defer func() {
		if s.cancel != nil {
			s.cancel()
		}
	}()

	err := s.d.Stop(ctx)
	if err != nil {
		return err
	}

	zap.L().Info("daemon server stop successfully")
	return nil
}
