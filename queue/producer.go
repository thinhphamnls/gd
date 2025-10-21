package queue

import (
	"fmt"
	"github.com/thinhphamnls/gd/logger"
	"time"

	"github.com/IBM/sarama"
	"go.uber.org/zap"

	"github.com/thinhphamnls/gd/config"
)

type IProducer interface {
	PushMessage(topic string, mgs string) error
	Close()
}

type producer struct {
	sugar          *zap.SugaredLogger
	producerClient sarama.SyncProducer
}

func NewProducer(cf bootstrap.Queue, zap logger.ILogger) (IProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	producerClient, err := sarama.NewSyncProducer(cf.Brokers, config)
	if err != nil {
		return nil, fmt.Errorf("producer client init fail, %s", err)
	}

	return producer{
		sugar:          zap.Get(),
		producerClient: producerClient,
	}, nil
}

func (p producer) PushMessage(topic string, mgs string) error {
	message := &sarama.ProducerMessage{
		Topic:     topic,
		Value:     sarama.StringEncoder(mgs),
		Timestamp: time.Now().UTC(),
	}

	partition, offset, err := p.producerClient.SendMessage(message)
	if err != nil {
		p.sugar.Warnf("partition: %d, offset: %d", partition, offset)
		return fmt.Errorf("send message to kafka failed, %v", err)
	}

	return nil
}

func (p producer) Close() {
	if err := p.producerClient.Close(); err != nil {
		return
	}
}
