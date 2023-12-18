package app

import (
	"context"
	"os"
	"time"

	"github.com/RaphaelL2e/golang-toolkit/delivery"
)

type Option func(o *options)

type options struct {
	id       string
	name     string
	version  string
	metadata map[string]string

	ctx  context.Context
	sigs []os.Signal

	stopTimeout time.Duration
	deliveries  []delivery.Delivery
}

func ID(id string) Option {
	return func(o *options) {
		o.id = id
	}
}

func Name(name string) Option {
	return func(o *options) {
		o.name = name
	}
}

func Version(version string) Option {
	return func(o *options) {
		o.version = version
	}
}

func Metadata(md map[string]string) Option {
	return func(o *options) {
		o.metadata = md
	}
}

func Deliveries(deliveries ...delivery.Delivery) Option {
	return func(o *options) {
		o.deliveries = deliveries
	}
}

func Signal(sigs ...os.Signal) Option {
	return func(o *options) {
		o.sigs = sigs
	}
}

func StopTimeout(t time.Duration) Option {
	return func(o *options) {
		o.stopTimeout = t
	}
}
