package app

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

// App is an application components lifecycle manager.
type App struct {
	opts   options
	ctx    context.Context
	cancel func()
}

// New create an application lifecycle manager.
func New(opts ...Option) *App {
	o := options{
		ctx:         context.Background(),
		sigs:        []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
		stopTimeout: 10 * time.Second,
	}

	for _, opt := range opts {
		opt(&o)
	}
	ctx, cancel := context.WithCancel(o.ctx)
	return &App{
		ctx:    ctx,
		cancel: cancel,
		opts:   o,
	}
}

// ID returns app instance id.
func (a *App) ID() string { return a.opts.id }

// Name returns service name.
func (a *App) Name() string { return a.opts.name }

// Version returns service version.
func (a *App) Version() string { return a.opts.version }

// Metadata returns service metadata.
func (a *App) Metadata() map[string]string { return a.opts.metadata }

// Run executes all OnStart hooks registered with the application's Lifecycle
func (a *App) Run() error {
	zap.L().Info("run app")

	eg, ctx := errgroup.WithContext(a.ctx)
	wg := sync.WaitGroup{}
	for _, delivery := range a.opts.deliveries {
		delivery := delivery
		eg.Go(func() error {
			<-ctx.Done()
			stopCtx, cancel := context.WithTimeout(a.opts.ctx, a.opts.stopTimeout)
			defer cancel()
			return delivery.Stop(stopCtx)
		})

		wg.Add(1)
		eg.Go(func() error {
			wg.Done()
			return delivery.Start(ctx)
		})
	}
	wg.Wait()

	c := make(chan os.Signal, 1)
	signal.Notify(c, a.opts.sigs...)
	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				zap.L().Info("receive signal")

				if err := a.Stop(); err != nil {
					zap.L().Error("failed to stop app", zap.Error(err))
					return err
				}
			}
		}
	})
	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

// Stop stops all OnStop hooks registered with the application's Lifecycle
func (a *App) Stop() error {
	zap.L().Info("stop app")
	if a.cancel != nil {
		a.cancel()
	}
	return nil
}
