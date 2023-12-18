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
