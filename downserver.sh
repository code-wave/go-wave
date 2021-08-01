#!/bin/bash

set -e
# if $1 is null
if [[ -z "$1" ]] || [[ "$1" = *"-"* ]]; then
    docker-compose down $1 $2 $3
    echo "All services down ..."
else
    # docker compose stop $1
    docker-compose rm --stop $1 
    echo "Service" $1 "down ..."
fi

# orphan containers remove ...
# if [[ ! -z "$(docker compose ps | grep running)" ]]; then
# 	echo "server remains orpahn containers, so remove orphan containers"
#     sleep 1s
#     echo -e "\nContainers Status"
#     docker compose ps
#     docker compose down -v --remove-orphans
# fi

sleep 0.5s
echo -e "\nContainers Status"
docker compose ps
