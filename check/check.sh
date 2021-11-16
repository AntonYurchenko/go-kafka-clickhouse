#!/bin/bash

echo "#### Send event to test topic #####"
echo '{"text":"Hello kafka!"}' | kcat -P -b localhost:9092 -t topic1

echo "#### Read event from test topic #####"
kcat -C -b localhost:9092 -t topic1 -c 1