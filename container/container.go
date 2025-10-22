package gdcontainer

import (
	"github.com/thinhphamnls/gd/config"
	"github.com/thinhphamnls/gd/logger"
)

func BuildDatabase(cf gdconfig.IConfig, zap gdlogger.ILogger) IDataBaseProvider {
	database, cleanup, err := NewDatabase(cf.GetDatabase(), zap)
	if err != nil {
		cleanup()
		zap.Get().Fatalf("init database failed: %v", err)
	}

	return database
}

func BuildRedis(cf gdconfig.IConfig, zap gdlogger.ILogger) IRedisProvider {
	redis, cleanup, err := NewRedis(cf.GetCache(), zap)
	if err != nil {
		cleanup()
		zap.Get().Fatalf("init redis failed: %v", err)
	}

	return redis
}

func BuildQueue(cf gdconfig.IConfig, zap gdlogger.ILogger) IProducer {
	queueClient, err := NewProducer(cf.GetQueue(), zap)
	if err != nil {
		zap.Get().Fatalf("init queue failed: %v", err)
	}

	return queueClient
}
