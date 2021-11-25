#!/bin/bash

JMETER_PATH='jmeter'

${JMETER_PATH} -n \
    -t ./rps_test.jmx \
    -l ./rps_test.csv \
    -e \
    -o ./report \
    -Jhost=0.0.0.0 \
    -Jport=8080 \
    -Jpath='/' \
    -Jusers=1 \
    -Jloop=1 \
    -Jramp=1