#!/bin/bash

go run . $1 || exit 10
chmod -R 775 build/ || exit 10
(cd build && go build . errors) || exit 10
(cd build/app && npm install && npm run build && npm run dev) || exit 10