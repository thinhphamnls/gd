package container

import (
	"github.com/thinhphamnls/gd/config"
	"github.com/thinhphamnls/gd/logger"
	"github.com/thinhphamnls/gd/queue"
)

type IBaseContainerProvider interface {
	RedisProvider() IRedisProvider
	DatabaseProvider() IDataBaseProvider
	QueueProvider() queue.IProducer
}

type BaseContainerProvider struct {
	redisProvider    IRedisProvider
	databaseProvider IDataBaseProvider
	queueProvider    queue.IProducer
}

func (p BaseContainerProvider) RedisProvider() IRedisProvider {
	return p.redisProvider
}

func (p BaseContainerProvider) DatabaseProvider() IDataBaseProvider {
	return p.databaseProvider
}

func (p BaseContainerProvider) QueueProvider() queue.IProducer {
	return p.queueProvider
}

func buildDatabase(cf gdconfig.Config, zap logger.ILogger) IDataBaseProvider {
	database, cleanup, err := NewDatabase(cf.Database, zap)
	if err != nil {
		cleanup()
		zap.Get().Fatalf("init database failed: %v", err)
	}

	return database
}

func buildRedis(cf gdconfig.Config, zap logger.ILogger) IRedisProvider {
	redis, cleanup, err := NewRedis(cf.Cache, zap)
	if err != nil {
		cleanup()
		zap.Get().Fatalf("init redis failed: %v", err)
	}

	return redis
}

func buildQueue(cf gdconfig.Config, zap logger.ILogger) queue.IProducer {
	queueClient, err := queue.NewProducer(cf.Queue, zap)
	if err != nil {
		zap.Get().Fatalf("init queue failed: %v", err)
	}

	return queueClient
}
