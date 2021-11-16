#!/bin/bash

set -a; source .env; set +a

echo "##### Initialisation of volumes #####"
mkdir -p ${ZOOKEEPER_PERSISTENCE} ${KAFKA_PERSISTENCE} ${CLICKHOUSE_PERSISTENCE} && echo "ok"

echo "##### Pulling docker images #####"
docker pull bitnami/zookeeper:${ZOOKEEPER_IMAGE_VERSION}
docker pull bitnami/kafka:${KAFKA_IMAGE_VERSION}

echo "##### Building go-entrypoint #####"
docker build \
    -f ./entrypoint/Dockerfile \
    -t github.com/antonyurchenko/entrypoint:${ENTRYPOINT_VERSION} \
    ./entrypoint

echo "##### Deploy all services on docker swarm #####"
docker stack deploy \
    --compose-file ./compose.yml \
    "${PROJECT_NAME}"
sleep 15

echo "##### Check status of all services #####"
docker stack services "${PROJECT_NAME}"