#!/bin/bash
cd vuetify-admin

cp ./src/config/index.js.prod ./src/config/index.js

echo "yarn install..."

yarn install --registry=https://registry.npm.taobao.org --chromedriver_cdnurl=https://npm.taobao.org/mirrors/chromedriver

echo "build..."

npm run build

echo "build complete"