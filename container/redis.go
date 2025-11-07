package gdcontainer

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"github.com/thinhphamnls/gd/config"
	"github.com/thinhphamnls/gd/logger"
)

const (
	redisMaxRetryBackoff = 30 * time.Second
	redisMinRetryBackoff = 1 * time.Second
	redisMaxRetries      = 10
)

type IRedisProvider interface {
	Get() *redis.Client
}

type redisProvider struct {
	redisClient *redis.Client
	sugar       *zap.SugaredLogger
}

func NewRedis(cf gdconfig.IBaseConfig, zap gdlogger.IBaseLogger) (IRedisProvider, func(), error) {
	var (
		data    = &redisProvider{sugar: zap.Get()}
		err     error
		cfRedis = cf.GetCache()
	)

	cleanup := func() {
		if data != nil && data.Get() != nil {
			_ = data.Get().Close()
		}

		zap.Get().Info("cleanup and close redis")
	}

	if cfRedis.Redis.Host != "" {
		data.redisClient, err = connectRedis(cfRedis.Redis)
		if err != nil {
			return nil, cleanup, err
		}
	}

	return data, cleanup, nil
}

func connectRedis(cfRedis gdconfig.Redis) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:            cfRedis.Host + ":" + cfRedis.Port,
		Password:        cfRedis.Password,
		DB:              cfRedis.Db,
		MinRetryBackoff: redisMinRetryBackoff,
		MaxRetryBackoff: redisMaxRetryBackoff,
		MaxRetries:      redisMaxRetries,
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
