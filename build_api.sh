#!/bin/bash
export GOPATH=$(pwd)/vuetify-admin-api

echo "GOPATH=$GOPATH"

cd vuetify-admin-api/src/vuetify-admin-api

echo "copy glide mirror..."

cp ./.glide/mirrors.yaml ~/.glide/mirrors.yaml

echo "glide install..."

glide install

echo "build..."
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o vuetify-admin-api ./app

chmod +x vuetify-admin-api

echo "build complete"