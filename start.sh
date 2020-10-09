#!/usr/bin/env bash
export GMAP_API=$1
docker-compose build
docker-compose up -d
docker-compose ps