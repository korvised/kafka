package services

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	events "github.com/korvised/kafka-events"
	"reflect"
)

type EventProducer interface {
	Produce(event events.Event) error
}

type eventProducer struct {
	producer sarama.SyncProducer
}

func NewEventProducer(producer sarama.SyncProducer) EventProducer {
	return &eventProducer{producer}
}

func (s *eventProducer) Produce(event events.Event) error {
	topic := reflect.TypeOf(event).Name()

	fmt.Println("Produce topic: ", topic)

	value, err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(value),
	}

	_, _, err = s.producer.SendMessage(msg)
	if err != nil {
		return err
	}

	return err
}
