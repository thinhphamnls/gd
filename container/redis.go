package container

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"sync-quickbooks-v3/config"
	"sync-quickbooks-v3/logger"
)

const (
	maxRetryBackoff = 30 * time.Second
	minRetryBackoff = 1 * time.Second
	maxRetries      = 10
)

type IRedisProvider interface {
	Get() *redis.Client
}

type redisProvider struct {
	redisClient *redis.Client
	sugar       *zap.SugaredLogger
}

func NewRedis(cf bootstrap.Cache, zap logger.ILogger) (IRedisProvider, func(), error) {
	var (
		data    = &redisProvider{sugar: zap.Get()}
		err     error
		cfRedis = cf.Redis
	)

	cleanup := func() {
		if data != nil && data.Get() != nil {
			_ = data.Get().Close()
		}

		zap.Get().Info("cleanup and close redis")
	}

	if cfRedis.Host != "" {
		data.redisClient, err = connectRedis(cfRedis)
		if err != nil {
			return nil, cleanup, err
		}
	}

	return data, cleanup, nil
}

func connectRedis(cfRedis bootstrap.Redis) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:            cfRedis.Host + ":" + cfRedis.Port,
		Password:        cfRedis.Password,
		DB:              cfRedis.Db,
		MinRetryBackoff: minRetryBackoff,
		MaxRetryBackoff: maxRetryBackoff,
		MaxRetries:      maxRetries,
		DialTimeout:     time.Duration(cfRedis.DialTimeout) * time.Second,
		ReadTimeout:     time.Duration(cfRedis.ReadTimeout) * time.Second,
		WriteTimeout:    time.Duration(cfRedis.WriteTimeout) * time.Second,
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return redisClient, nil
}

func (p *redisProvider) Get() *redis.Client {
	return p.redisClient
}
