package redis

import (
	"fmt"
	"github.com/RaphaelL2e/golang-toolkit/constant"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"strings"
)

type DistributeLocker struct {
	namespace string

	rs *redsync.Redsync
}

func NewDistributeLocker(client *RedisClient) DistributeLocker {
	pool := goredis.NewPool(client)
	rs := redsync.New(pool)
	return DistributeLocker{
		namespace: client.namespace,
		rs:        rs,
	}
}

func (d *DistributeLocker) Mutex(args ...string) *redsync.Mutex {
	return d.rs.NewMutex(
		d.getName(args...),
		redsync.WithTries(constant.RedsyncTries),
		redsync.WithExpiry(constant.RedsyncExpiry),
	)
}

func (d *DistributeLocker) MutexWithTries(tires int, args ...string) *redsync.Mutex {
	return d.rs.NewMutex(
		d.getName(args...),
		redsync.WithTries(tires),
		redsync.WithExpiry(constant.RedsyncExpiry),
	)
}

func (d *DistributeLocker) getName(args ...string) string {
	return fmt.Sprintf("%s{%s}%slock",
		d.namespace,
		strings.Join(args, RedisDilimiters[0]),
		RedisDilimiters[1],
	)
}
