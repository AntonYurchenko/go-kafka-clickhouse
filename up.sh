#!/bin/bash

set -a; source .env; set +a

echo "##### Pulling necessary docker images #####"
docker pull bitnami/zookeeper:${ZOOKEEPER_IMAGE_VERSION}
docker pull bitnami/kafka:${KAFKA_IMAGE_VERSION}

echo "##### Deploy all services on docker swarm #####"
docker stack deploy \
    --compose-file ./compose.yml \
    "${PROJECT_NAME}"
sleep 5

echo "##### Check status of all services #####"
docker stack services "${PROJECT_NAME}"

echo "#### Send event to test topic #####"
echo '{"event":"Hello kafka!"}' | kcat -P -b localhost:9092 -t test_topic

echo "#### Read event from test topic #####"
kcat -C -b localhost:9092 -t test_topic -c 1