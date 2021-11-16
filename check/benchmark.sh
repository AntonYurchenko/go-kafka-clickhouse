#!/bin/bash

ab -k -n 100000 -c 10 \
    -T 'application/json' \
    -p ./data.json \
    http://0.0.0.0:8080/