package redis

import (
	"context"
	"strings"
	"time"

	redis "github.com/go-redis/redis/v8"

	"github.com/renjingneng/a_simple_go_project/core/container"
)

type Base interface {
	Get(key string) (string, error)
	Set(key string, value interface{}, expiration time.Duration) (string, error)
}
type singleBase struct {
	Dbname string
	Mode   string
	ctx    context.Context
	Dbptr  *redis.Client
}
type clusterBase struct {
	Dbname string
	Mode   string
	ctx    context.Context
	Dbptr  *redis.ClusterClient
}

func NewBase(Dbname string, Mode string) Base {
	if strings.ToUpper(Mode) == "SINGLE" {
		dbptr := container.GetEntityFromRedisContainer(Dbname, "Single")
		db := &singleBase{
			Dbname: Dbname,
			Mode:   "Single",
			ctx:    context.Background(),
		}
		db.Dbptr, _ = dbptr.(*redis.Client)
		return db
	} else if strings.ToUpper(Mode) == "CLUSTER" {
		dbptr := container.GetEntityFromRedisContainer(Dbname, "Cluster")
		db := &clusterBase{
			Dbname: Dbname,
			Mode:   "Cluster",
			ctx:    context.Background(),
		}
		db.Dbptr, _ = dbptr.(*redis.ClusterClient)
		return db
	} else {
		return nil
	}
}

func (b *singleBase) Get(key string) (string, error) {
	val, err := b.Dbptr.Get(b.ctx, key).Result()
	if err != nil {
		return val, err
	} else {
		return val, nil
	}
}
func (b *singleBase) Set(key string, value interface{}, expiration time.Duration) (string, error) {
	val, err := b.Dbptr.Set(b.ctx, key, value, expiration).Result()
	if err != nil {
		return val, err
	} else {
		return val, nil
	}
}

func (b *clusterBase) Get(key string) (string, error) {
	val, err := b.Dbptr.Get(b.ctx, key).Result()
	if err != nil {
		return val, err
	} else {
		return val, nil
	}
}

func (b *clusterBase) Set(key string, value interface{}, expiration time.Duration) (string, error) {
	val, err := b.Dbptr.Set(b.ctx, key, value, expiration).Result()
	if err != nil {
		return val, err
	} else {
		return val, nil
	}
}
