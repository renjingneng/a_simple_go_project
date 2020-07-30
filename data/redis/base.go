package redis

import (
	"context"
	"time"

	redis "github.com/go-redis/redis/v8"

	"github.com/renjingneng/a_simple_go_project/core/container"
)

type Base struct {
	Dbname string
	Mode   string
	ctx    context.Context
	Dbptr  *redis.Client
}

func NewBase(Dbname string, Mode string) *Base {
	return &Base{
		Dbname: Dbname,
		Mode:   Mode,
		ctx:    context.Background(),
		Dbptr:  container.GetEntityFromRedisContainer(Dbname, Mode),
	}
}

func (b *Base) Get(key string) (string, error) {
	val, err := b.Dbptr.Get(b.ctx, key).Result()
	if err != nil {
		return val, err
	} else {
		return val, nil
	}
}

func (b *Base) Set(key string, value interface{}, expiration time.Duration) (string, error) {
	val, err := b.Dbptr.Set(b.ctx, key, value, expiration).Result()
	if err != nil {
		return val, err
	} else {
		return val, nil
	}
}
