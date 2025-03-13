#!/bin/bash

# Start containers
docker-compose -f environment/redis-cluster/docker-compose-dev.yml up
echo "[TIPsGO]: vetautet server start ..."