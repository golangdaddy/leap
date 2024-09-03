#!/bin/bash

go mod tidy

GONOPROXY=github.com/golangdaddy go get -u github.com/golangdaddy/leap

go test || exit 10
chmod -R 775 build/ || exit 10

echo "ENVIRONMENT=dev" > build/app/.env.local
echo "HANDCASH_APP_ID=${HANDCASH_APP_ID}" >> build/app/.env.local
echo "HANDCASH_APP_SECRET=${HANDCASH_APP_SECRET}" >> build/app/.env.local

(cd build/app && npm install && npm run dev) || exit 10