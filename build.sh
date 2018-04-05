#!/bin/bash
# git reset --hard HEAD
# git pull && git submodule update --init --recursive
sh ./build_api.sh
sh ./build_web.sh

docker-compose -f docker-compose.yml up -d --build 

docker restart vuetify-admin
docker restart vuetify-admin-api




