package gdcontainer

import (
	"fmt"
	"github.com/google/uuid"
	"time"

	"github.com/IBM/sarama"
	"go.uber.org/zap"

	"github.com/thinhphamnls/gd/config"
	"github.com/thinhphamnls/gd/logger"
)

type IProducer interface {
	PushMessage(topic string, mgs string) error
	Close()
}

type producer struct {
	sugar          *zap.SugaredLogger
	producerClient sarama.SyncProducer
}

func NewProducer(cf gdconfig.IBaseConfig, zap gdlogger.IBaseLogger, cfKfk *sarama.Config) (IProducer, error) {
	producerClient, err := sarama.NewSyncProducer(cf.GetQueue().Brokers, cfKfk)
	if err != nil {
		return nil, fmt.Errorf("producer client init failed: %s", err)
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
		Headers: []sarama.RecordHeader{
			{Key: []byte("message_id"), Value: []byte(uuid.NewString())},
		},
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
