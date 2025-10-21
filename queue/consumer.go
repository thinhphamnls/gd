package queue

import (
	"context"
	"fmt"
	"github.com/thinhphamnls/gd/config"

	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

type IConsumer interface {
	Start(ctx context.Context, handler *ConsumerHandler) func()
	Stop() func()
}

type consumer struct {
	sugar         *zap.SugaredLogger
	consumerGroup sarama.ConsumerGroup
	topics        []string
}

func NewConsumer(cf bootstrap.Queue, sugar *zap.SugaredLogger) (IConsumer, error) {
	if len(cf.Brokers) == 0 {
		return nil, fmt.Errorf("broker list is empty")
	}

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Offsets.AutoCommit.Enable = false

	consumerClient, err := sarama.NewConsumerGroup(cf.Brokers, cf.GroupId, config)
	if err != nil {
		return nil, fmt.Errorf("new consumer group failed:%v", err)
	}

	return &consumer{
		sugar:         sugar,
		consumerGroup: consumerClient,
		topics:        []string{cf.Topic},
	}, nil
}

func (c consumer) Start(ctx context.Context, handler *ConsumerHandler) func() {
	return func() {
		go func(ctx context.Context, topics []string, handler *ConsumerHandler) {
			for {
				if err := c.consumerGroup.Consume(ctx, topics, handler); err != nil {
					c.sugar.Errorf("consuming messages failed, %v", err)
				}
			}
		}(ctx, c.topics, handler)
	}
}

func (c consumer) Stop() func() {
	return func() {
		if err := c.consumerGroup.Close(); err != nil {
			c.sugar.Errorf("closing consumer failed, %v", err)
		}
	}
}
