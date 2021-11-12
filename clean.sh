#!/bin/bash

set -a; source .env; set +a

echo "##### Cleaning of volumes #####"
rm -rf ${ZOOKEEPER_PERSISTENCE} ${KAFKA_PERSISTENCE} ${CLICKHOUSE_PERSISTENCE} && echo "ok"