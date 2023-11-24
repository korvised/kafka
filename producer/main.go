package main

import (
	"fmt"
	"github.com/IBM/sarama"
)

func main() {
	servers := []string{"localhost:9092"}

	producer, err := sarama.NewSyncProducer(servers, nil)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	fmt.Println("")

	msg := sarama.ProducerMessage{
		Topic: "longhello",
		Value: sarama.StringEncoder("hello long"),
	}

	partition, offset, err := producer.SendMessage(&msg)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Partition(%d) | Offset(%d) \n", partition, offset)
}
