package main

import (
	"flag"
	"os"
	"os/signal"
	"strings"
)

var (
	entrypoint = flag.String("e", "localhost:8080", "Endpoint of http server")
	brokers    = flag.String("b", "localhost:9092", "Kafka brokers split by comma")
	topic      = flag.String("t", "topic1", "Kafka target topic")
	cache      = flag.Uint("c", 1, "Size of internal cache")
)

func main() {

	flag.Parse()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	events := make(chan Event, *cache)
	go reciever(events, *entrypoint)
	go sender(events, strings.Split(*brokers, ","), *topic)

	<-signals
}
