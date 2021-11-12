package main

import (
	"encoding/json"
	"fmt"

	"github.com/Shopify/sarama"
)

func sender(events <-chan Event, brokers []string, topic string) {

	provider, err := sarama.NewAsyncProducer(brokers, nil)
	if err != nil {
		panic(err)
	}
	defer provider.Close()

	for {
		select {
		case event := <-events:
			value, _ := json.Marshal(event)
			provider.Input() <- &sarama.ProducerMessage{
				Topic: topic,
				Key:   nil,
				Value: sarama.ByteEncoder(value),
			}
		case err := <-provider.Errors():
			fmt.Printf("Error of sending: %v", err)
		}
	}

}
