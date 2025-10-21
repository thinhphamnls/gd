package queue

import "github.com/IBM/sarama"

type MessageConsumer interface {
	ConsumerMessage(msg *sarama.ConsumerMessage, session sarama.ConsumerGroupSession)
}

type ConsumerHandler struct {
	consumer MessageConsumer
}

func NewConsumerGroupHandler(consumer MessageConsumer) *ConsumerHandler {
	return &ConsumerHandler{
		consumer: consumer,
	}
}

func (c *ConsumerHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c *ConsumerHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c *ConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		c.consumer.ConsumerMessage(msg, session)
	}
	return nil
}
