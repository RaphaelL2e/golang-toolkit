package delivery

import "context"

type Delivery interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}
