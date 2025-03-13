#!/bin/bash

# Start containers
docker-compose -f environment/docker-compose-dev.yml down
echo "[TIPsGO]: vetautet server stop ..."