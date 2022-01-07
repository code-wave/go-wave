#!/bin/bash

set -e

if [[ "$1" = "--local" ]]; then
    echo "Run $2 server for local mode..."
    docker-compose -f docker-compose.yml -f docker-compose.local.yml up -d --build $2
elif [[ "$1" = "--dev" ]]; then
    echo "Run $2 server for develop mode..."
    docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d --build $2
else
    echo "Usage: bash runserver.sh [--dev | --prod]"
fi


sleep 0.5s
echo -e "\nContainers Status"
docker-compose ps
