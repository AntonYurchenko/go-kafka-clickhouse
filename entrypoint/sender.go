package main

import (
	"encoding/json"
	"fmt"

	"github.com/Shopify/sarama"
)

// sender is a funcrion for writting events to selected Kafka topic.
func sender(events <-chan *Event, brokers []string, topic string) {

	// start Kafka provider.
	provider, err := sarama.NewAsyncProducer(brokers, nil)
	if err != nil {
		panic(err) //todo: processing error is here
	}
	defer provider.Close()

	// writting events.
	for {
		select {
		case event := <-events:
			value, _ := json.Marshal(*event)
			provider.Input() <- &sarama.ProducerMessage{
				Topic: topic,
				Key:   nil,
				Value: sarama.ByteEncoder(value),
			}
		case err := <-provider.Errors():
			fmt.Printf("Error of sending: %v", err) //todo: processing error is here
		}
	}

}
