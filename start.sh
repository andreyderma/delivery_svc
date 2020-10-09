#!/usr/bin/env bash
GMAP_API=$1
sed -i -e "s~GMAP_API~$GMAP_API~g" docker-compose.yml
#docker-compose build
docker-compose up -d
docker-compose ps