package services

import (
	"github.com/IBM/sarama"
)

type consumerHandler struct {
	eventHandler EventHandler
}

func NewConsumerHandler(eventHandler EventHandler) sarama.ConsumerGroupHandler {
	return &consumerHandler{eventHandler}
}

func (c *consumerHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c *consumerHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c *consumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		c.eventHandler.Handle(msg.Topic, msg.Value)
		session.MarkMessage(msg, "") // marked for had received this message
	}

	return nil
}
