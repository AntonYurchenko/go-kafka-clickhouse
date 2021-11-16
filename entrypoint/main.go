package main

import (
	"flag"
	"os"
	"os/signal"
	"runtime"
	"strings"
)

var (
	entrypoint = flag.String("e", "0.0.0.0:8080", "Entrypoint of http server")
	brokers    = flag.String("b", "0.0.0.0:9092", "Kafka brokers split by comma")
	topic      = flag.String("t", "topic1", "Kafka target topic")
	cache      = flag.Uint("c", 100, "Size of internal cache")
)

func main() {

	flag.Parse()

	signals := make(chan os.Signal, 1)
	events := make(chan Event, *cache)
	cpuCount := runtime.NumCPU()

	signal.Notify(signals, os.Interrupt)
	runtime.GOMAXPROCS(cpuCount)

	go reciever(events, *entrypoint)
	for ; cpuCount > 0; cpuCount-- {
		go sender(events, strings.Split(*brokers, ","), *topic)
	}

	<-signals
}
