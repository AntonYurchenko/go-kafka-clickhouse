# Events collector with http interface

## Architectural diagram
```text
                              + ------------ +
                         + -> | entrypoint 1 | - +
                         |    + ------------ +   |
http                     |                       |       _____________
port  + ------------ +   |    + ------------ +   |     /\              \
   >- | docker swarm | - + -> | entrypoint 2 | - + -> (  ) Apache Kafka )
      |   balancer   |   |    + ------------ +   |     \/______________/
      + ------------ +   |                       |            ||
                         |    + ------------ +   |            \/
                         + -> | entrypoint 3 | - +        __________   
                              + ------------ +          ( __________ )
                                                        |            |
                                                        | ClickHouse |
                                                        | __________ |
```

## Start of all environment
Clone of this repository
```bash
git clone https://github.com/AntonYurchenko/go-kafka-clickhouse.git go-kafka-clickhouse
cd go-kafka-clickhouse
```

You can change some options in file `.env`: 
* name of a project 
* path to persistent data of containers 
* count of replices http endpoint
```text
PROJECT_NAME=go-kafka-clickhouse
ENTRYPOINT_REPLICAS=3
ZOOKEEPER_PERSISTENCE=./volumes/zookeeper
KAFKA_PERSISTENCE=./volumes/kafka
CLICKHOUSE_PERSISTENCE=./volumes/clickhouse/data
```

Up of all services
```bash
./up.sh
```