#!/bin/bash

set -a; source .env; set +a

docker stack rm "${PROJECT_NAME}"