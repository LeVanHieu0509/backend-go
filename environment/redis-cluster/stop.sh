#!/bin/bash

# Start containers
docker-compose -f environment/redis-cluster/docker-compose-dev.yml down
echo "[TIPsGO]: vetautet server stop ..."