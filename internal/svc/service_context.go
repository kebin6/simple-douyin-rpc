package svc

import (
	"context"
	"github.com/kebin6/simple-douyin-rpc/ent"
	_ "github.com/kebin6/simple-douyin-rpc/ent/runtime"
	"github.com/kebin6/simple-douyin-rpc/internal/config"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
	Config config.Config
	DB     *ent.Client
	Redis  redis.UniversalClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := ent.NewClient(
		ent.Log(logx.Info), // logger
		ent.Driver(c.DatabaseConf.NewNoCacheDriver()),
		ent.Debug(), // debug mode
	)

	return &ServiceContext{
		Config: c,
		DB:     db,
		Redis:  c.RedisConf.MustNewUniversalRedis(),
	}
}

type RedisCache struct {
	Ctx   context.Context
	Redis redis.UniversalClient
}

func (r *RedisCache) Get(key string) interface{} {
	res, err := r.Redis.Get(r.Ctx, key).Result()
	if err != nil {
		return nil
	}
	return res
}

func (r *RedisCache) Set(key string, val interface{}, timeout time.Duration) error {
	_, err := r.Redis.Set(r.Ctx, key, val, timeout).Result()
	return err
}

func (r *RedisCache) IsExist(key string) bool {
	res, err := r.Redis.Exists(r.Ctx, key).Result()
	if err != nil {
		return false
	}
	return res > 0
}

func (r *RedisCache) Delete(key string) error {
	return r.Redis.Del(r.Ctx, key).Err()
}

func (l *ServiceContext) NewRedisCache(ctx context.Context, redis redis.UniversalClient) *RedisCache {
	return &RedisCache{
		Ctx:   ctx,
		Redis: redis,
	}
}
