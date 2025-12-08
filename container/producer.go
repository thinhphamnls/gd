package gdcontainer

import (
	"fmt"
	"time"

	"github.com/IBM/sarama"
	"github.com/google/uuid"
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

func NewProducer(cf gdconfig.BaseConfig, zap gdlogger.IBaseLogger, cfKfk *sarama.Config) (IProducer, error) {
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
	messageId := uuid.NewString()
	message := &sarama.ProducerMessage{
		Topic:     topic,
		Value:     sarama.StringEncoder(mgs),
		Timestamp: time.Now().UTC(),
		Headers: []sarama.RecordHeader{
			{Key: []byte("message_id"), Value: []byte(messageId)},
		},
	}

	partition, offset, err := p.producerClient.SendMessage(message)
	if err != nil {
		return fmt.Errorf("send message to kafka failed, %v", err)
	}

	p.sugar.Infof("Message published MessageId: %v Partition: %d Offset: %d", messageId, partition, offset)
	return nil
}

func (p producer) Close() {
	if err := p.producerClient.Close(); err != nil {
		return
	}
}
