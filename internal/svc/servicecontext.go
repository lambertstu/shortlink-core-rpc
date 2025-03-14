package svc

import (
	"github.com/lambertstu/shortlink-core-rpc/internal/config"
	shortlink "github.com/lambertstu/shortlink-core-rpc/mongo/shortlink"
	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

const (
	bloomKey = "redis-bloom:shortlink"
	bloomBit = 64
)

type ServiceContext struct {
	Config         config.Config
	Redis          *redis.Redis
	BloomFilter    *bloom.Filter
	ShortlinkModel shortlink.ShortlinkModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	rds, filter := getRedisInstance(c)
	return &ServiceContext{
		Config:         c,
		Redis:          rds,
		BloomFilter:    filter,
		ShortlinkModel: shortlink.NewShortlinkModel(c.MongoUrl, shortlink.DB, shortlink.Collection),
	}
}

func getRedisInstance(c config.Config) (*redis.Redis, *bloom.Filter) {
	rds := redis.MustNewRedis(c.RedisConf)
	filter := bloom.New(rds, bloomKey, bloomBit)
	return rds, filter
}
